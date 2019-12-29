package cert

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"net"
	"time"
)

// Client represents a client certificate.
type Client struct {
	*Certificate
	CommonName string
	ips        []net.IP
	dns        []string
}

// GenCert generates a certificate.
func (cert *Client) GenCert() error {
	cert.cert = &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject: pkix.Name{
			CommonName: cert.CommonName,
		},
		IPAddresses:           cert.ips,
		DNSNames:              cert.dns,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	var err error
	cert.certBytes, err = x509.CreateCertificate(rand.Reader, cert.cert, cert.cert, cert.pubKey, cert.privKey)
	if err != nil {
		return err
	}

	return err
}

// NewClientCert returns an instance of a client certificate with common name
// and IPs/DNS names as SAN.
func NewClientCert(commonName string, ips, dns []string) *Client {
	cert := Client{
		CommonName:  commonName,
		Certificate: &Certificate{},
		dns:         dns,
	}

	for _, ip := range ips {
		cert.ips = append(cert.ips, net.ParseIP(ip))
	}

	return &cert
}
