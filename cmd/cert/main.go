package main

import (
	"errors"
	"flag"
	"github.com/marcboeker/tentacle/cert"
	"log"
	"strings"
)

var (
	action, organization, commonName, ips, dns *string
	caCertFile, caKeyFile, certFile, keyFile   *string
)

func init() {
	action = flag.String("action", "", "Action to to (ca - create new CA, sign - sign client cert)")
	organization = flag.String("organization", "", "CA organization name")
	commonName = flag.String("common-name", "", "common name for client cert")
	ips = flag.String("ips", "", "List of comma separated IP addresses")
	dns = flag.String("dns", "", "List of comma separated DNS names")
	caCertFile = flag.String("ca-cert", "ca.crt", "Name of certificate file")
	caKeyFile = flag.String("ca-key", "ca.key", "Name of key file")
	certFile = flag.String("cert", "node.crt", "Name of certificate file")
	keyFile = flag.String("key", "node.key", "Name of key file")
	flag.Parse()
}

func main() {
	if err := validateArgs(*action, *caCertFile, *caKeyFile, *certFile, *keyFile); err != nil {
		log.Fatal(err)
	}

	switch *action {
	case "ca":
		err := genCA(*organization, *caCertFile, *caKeyFile)
		if err != nil {
			log.Fatal(err)
		}
	case "sign":
		err := genNode(*commonName, *caCertFile, *caKeyFile, *certFile, *keyFile, *ips, *dns)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func validateArgs(action, caCertFile, caKeyFile, certFile, keyFile string) error {
	if len(action) == 0 {
		return errors.New("no action specified")
	}

	if len(certFile) == 0 {
		return errors.New("no cert file specified")
	}

	if len(keyFile) == 0 {
		return errors.New("no key file specified")
	}

	if len(caCertFile) == 0 {
		return errors.New("no CA cert file specified")
	}

	if len(caKeyFile) == 0 {
		return errors.New("no CA key file specified")
	}

	return nil
}

func genCA(organization, caCertFile, caKeyFile string) error {
	if len(organization) == 0 {
		log.Fatal("Please specify an organization")
	}

	ca := cert.NewCACert(organization)
	if err := ca.GenKeys(); err != nil {
		return err
	}
	if err := ca.GenCert(); err != nil {
		return err
	}
	if err := ca.WriteCert(caCertFile); err != nil {
		return err
	}
	err := ca.WriteKey(caKeyFile)
	return err
}

func genNode(commonName, caCertFile, caKeyFile, certFile, keyFile, ips, dns string) error {
	if len(commonName) == 0 {
		log.Fatal("Please specify a common name")
	}

	c, err := cert.Parse(caCertFile, caKeyFile)
	if err != nil {
		return err
	}

	ca := cert.CA{Certificate: c}
	cc := cert.NewClientCert(commonName, strings.Split(ips, ","), strings.Split(dns, ","))
	if err := cc.GenKeys(); err != nil {
		return err
	}
	if err := cc.GenCert(); err != nil {
		return err
	}
	if err := ca.Sign(cc); err != nil {
		return err
	}
	if err := cc.WriteCert(certFile); err != nil {
		return err
	}
	err = cc.WriteKey(keyFile)
	return err
}
