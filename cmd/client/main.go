package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"

	backoff "github.com/cenkalti/backoff/v3"

	"github.com/marcboeker/tentacle/client"
)

var (
	token, serverAddr, localAddr, certsDir *string
)

func init() {
	token = flag.String("token", "", "auth token")
	serverAddr = flag.String("server", "server:1337", "Tentacle server address ({IP/hostname}:{port})")
	localAddr = flag.String("local", ":0", "Bind to specific local instead of default address")
	certsDir = flag.String("certs-dir", "", "Directory containing ca.crt, node.crt and node.key")
	flag.Parse()
}

func main() {
	if err := validateArgs(*token, *certsDir); err != nil {
		log.Fatal(err)
	}

	transport, err := client.NewTransport(*localAddr, *certsDir)
	if err != nil {
		log.Fatal(err)
	}

	rpc, err := client.NewRPC(*token, *serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Get local listening port and register.
	if err := rpc.Register(transport.ListenPort()); err != nil {
		log.Fatal(err)
	}

	go func() {
		bo := backoff.NewExponentialBackOff()
		backOffFunc := func() error {
			return rpc.WaitForConnection(func(peerAddr string) error {
				return transport.PunchHole(peerAddr)
			})
		}
		err := backoff.Retry(backOffFunc, bo)
		if err != nil {
			log.Fatal("Could not connect to server")
			return
		}
	}()

	iface, err := client.SetupInterface()
	if err != nil {
		log.Fatal(err)
	}

	subnets, err := rpc.Subnets()
	if err != nil {
		log.Fatal(err)
	}

	// Add IPs and routes for subnets.
	var sn client.SubnetList
	for _, s := range subnets {
		subnet, err := client.NewSubnet(iface, s.Name, s.Ip, s.Cidr)
		if err != nil {
			log.Fatal(err)
		}
		sn.Add(subnet)
	}

	inbound := make(chan client.Packet)
	outbound := make(chan client.Packet)

	tw := client.NewTransportWrapper(rpc, transport, inbound)
	r := client.NewRouter(inbound, outbound, sn.LocalIPs(), sn.LocalNets())

	go iface.Read(inbound)
	go tw.Serve(*localAddr)
	go r.Route()

	for {
		select {
		case pkt := <-outbound:
			if pkt.DstType == client.DestinationLocal {
				if _, err := iface.Write(pkt.Data); err != nil {
					continue
				}
			}
			if pkt.DstType == client.DestinationRemote {
				stream, err := tw.GetOrConnect(pkt.DstAddr)
				if err != nil {
					// Drop packet if no connection can be made.
					continue
				}

				if err := stream.Write(pkt); err != nil {
					// Drop packet if stream is down.
					continue
				}
			}
		}
	}
}

func validateArgs(token, certsDir string) error {
	if len(token) == 0 {
		return errInvalidToken
	}

	if len(certsDir) == 0 {
		return errInvalidCertsDir
	}

	caFile := filepath.Join(certsDir, client.CAName)
	certFile := filepath.Join(certsDir, client.CertName)
	keyFile := filepath.Join(certsDir, client.KeyName)

	if _, err := os.Stat(caFile); os.IsNotExist(err) {
		return errMissingCACertFile
	}

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		return errMissingCertFile
	}

	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		return errMissingKeyFile
	}

	return nil
}

var (
	errInvalidToken      = errors.New("please specify a valid token")
	errInvalidCertsDir   = errors.New("could not find certs dir")
	errMissingCACertFile = errors.New("could not find ca.crt in certs dir")
	errMissingCertFile   = errors.New("could not find node.crt in certs dir")
	errMissingKeyFile    = errors.New("could not find node.key in certs dir")
)
