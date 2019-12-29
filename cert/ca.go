package cert

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

// NewCACert returns an instance of CA with the given organication name.
func NewCACert(organization string) *CA {
	return &CA{Organization: organization, Certificate: &Certificate{}}
}

// CA represents a CA certificate.
type CA struct {
	*Certificate
	Organization string
}

// GenCert generates a X509 CA certificate.
func (ca *CA) GenCert() error {
	ca.cert = &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject: pkix.Name{
			Organization: []string{ca.Organization},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	var err error
	ca.certBytes, err = x509.CreateCertificate(rand.Reader, ca.cert, ca.cert, ca.pubKey, ca.privKey)
	if err != nil {
		return err
	}

	return err
}

// Sign signs a given client certificate with the current CA.
func (ca CA) Sign(cert *Client) error {
	var err error
	cert.cert.Issuer = ca.cert.Subject
	cert.certBytes, err = x509.CreateCertificate(rand.Reader, cert.cert, ca.cert, cert.pubKey, ca.privKey)
	return err
}
