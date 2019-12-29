package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net"
	"path/filepath"
	"time"

	"github.com/lucas-clemente/quic-go"
)

const (
	// CAName is the file name of the CA cert.
	CAName = "ca.crt"
	// CertName is the file name of the node cert.
	CertName = "node.crt"
	// KeyName is the file name of the node private key.
	KeyName = "node.key"
)

var (
	// The basic config params for the QUIC transport.
	quicConf = quic.Config{
		HandshakeTimeout: time.Second * 10,
		KeepAlive:        true,
	}
)

const (
	// QUIC protocol version
	protoVer = "tentacle-p2p"
)

// Transport holds the predefined connection and its TLS certs.
type Transport struct {
	conn       *net.UDPConn
	cert       tls.Certificate
	caCertPool *x509.CertPool
}

// NewTransport returns an instance of the transport.
func NewTransport(localAddr, certsDir string) (*Transport, error) {
	cert, caCertPool, err := loadTLS(certsDir)
	if err != nil {
		return nil, err
	}

	addr, err := net.ResolveUDPAddr("udp4", localAddr)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		return nil, err
	}

	return &Transport{
		conn:       conn,
		cert:       cert,
		caCertPool: caCertPool,
	}, nil
}

// Server initializes a QUIC UDP server that accepts remote connections.
func (t *Transport) Server(handler func(stream *Stream)) error {
	tlsConf := &tls.Config{
		RootCAs:      t.caCertPool,
		Certificates: []tls.Certificate{t.cert},
		NextProtos:   []string{protoVer},
		ClientAuth:   tls.RequireAnyClientCert,
	}

	listener, err := quic.Listen(t.conn, tlsConf, &quicConf)
	if err != nil {
		return err
	}

	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			continue
		}

		stream, err := sess.AcceptStream(context.Background())
		if err != nil {
			continue
		}

		go handler(&Stream{Stream: stream, sess: sess})
	}
}

// Dial connects to the remote peer and performs a handshake.
func (t *Transport) Dial(peerAddress, localPeerAddr string) (*Stream, error) {
	tlsConf := &tls.Config{
		RootCAs:            t.caCertPool,
		Certificates:       []tls.Certificate{t.cert},
		InsecureSkipVerify: false,
		ServerName:         localPeerAddr,
		NextProtos:         []string{protoVer},
	}

	addr, err := net.ResolveUDPAddr("udp", peerAddress)
	if err != nil {
		return nil, err
	}

	sess, err := quic.Dial(t.conn, addr, "", tlsConf, &quicConf)
	if err != nil {
		return nil, err
	}

	stream, err := sess.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}

	return &Stream{Stream: stream, sess: sess}, nil
}

// ListenPort returns the local port the listener is listening on,
func (t *Transport) ListenPort() string {
	_, port, _ := net.SplitHostPort(t.conn.LocalAddr().String())
	return port
}

// PunchHole sends some UDP packets to the given UDP address to punch
// a hole in NAT firewalls.
func (t *Transport) PunchHole(peerAddr string) error {
	addr, err := net.ResolveUDPAddr("udp", peerAddr)
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		t.conn.WriteToUDP(nil, addr)
		time.Sleep(time.Second)
	}

	return nil
}

func loadTLS(certsDir string) (tls.Certificate, *x509.CertPool, error) {
	caFile := filepath.Join(certsDir, CAName)
	certFile := filepath.Join(certsDir, CertName)
	keyFile := filepath.Join(certsDir, KeyName)

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return tls.Certificate{}, nil, err
	}

	clientCACert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return tls.Certificate{}, nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(clientCACert)

	return cert, caCertPool, nil
}

// Stream wraps a QUIC stream.
type Stream struct {
	quic.Stream
	sess   quic.Session
	Closed bool
}

// Read reads a packet from the stream.
// If the read fails it closes the stream and session.
func (s *Stream) Read(cb func(pkt []byte, s *Stream)) error {
	if s.Closed {
		return errStreamClosed
	}

	buf := make([]byte, 2048)
	for {
		n, err := s.Stream.Read(buf)
		if err != nil {
			s.Stream.Close()
			s.sess.Close()
			s.Closed = true
			return err
		}

		cb(buf[:n], s)
	}
}

// Write writes a packet to the stream.
// If the write fails it closes the stream and session.
func (s *Stream) Write(pkt Packet) error {
	if s.Closed {
		return errStreamClosed
	}

	if _, err := s.Stream.Write(pkt.Data); err != nil {
		s.Stream.Close()
		s.sess.Close()
		s.Closed = true
		return err
	}

	return nil
}

var (
	errStreamClosed = errors.New("stream is already closed")
)
