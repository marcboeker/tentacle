package cert

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
)

// Certificate represents a basic certificate with its private and public key.
type Certificate struct {
	CommonName string
	privKey    ed25519.PrivateKey
	pubKey     ed25519.PublicKey
	cert       *x509.Certificate
	certBytes  []byte
}

// GenKeys generates a private and public key pair.
func (cert *Certificate) GenKeys() error {
	var err error
	cert.pubKey, cert.privKey, err = ed25519.GenerateKey(rand.Reader)
	return err
}

// WriteKey writes the private key to the given file name.
func (cert Certificate) WriteKey(out string) error {
	privBytes, err := x509.MarshalPKCS8PrivateKey(cert.privKey)
	if err != nil {
		return err
	}
	fh, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	return pem.Encode(fh, &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privBytes,
	})
}

// WriteCert writes the certificate to the given file name.
func (cert Certificate) WriteCert(out string) error {
	fh, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	return pem.Encode(fh, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.certBytes,
	})
}

// Parse parses a certificate with its private key from a file.
func Parse(certFile, keyFile string) (*Certificate, error) {
	cert := Certificate{}

	var err error
	cert.certBytes, err = ioutil.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	pemBlock, _ := pem.Decode(cert.certBytes)

	cert.cert, err = x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	keyBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	keyBlock, _ := pem.Decode(keyBytes)

	privKey, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, err
	}

	cert.privKey, _ = privKey.(ed25519.PrivateKey)
	cert.pubKey = cert.privKey.Public().(ed25519.PublicKey)

	return &cert, nil
}
