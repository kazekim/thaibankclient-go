/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package thaibankclient

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

const (
	x509CertificateCommonName = "ThaiBankClientGo"
	x509CertificateOrganization = "Jirawat.Kim"
)

type TLSCertGenerator struct {
	key *rsa.PrivateKey
}

func NewTLSCertGenerator() (*TLSCertGenerator, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("private key cannot be created: %v", err.Error())
	}

	return &TLSCertGenerator{
		key: key,
	}, nil
}

func (g *TLSCertGenerator) GeneratePEMPrivateKey() (*[]byte, error) {

	// Generate a pem block with the private key
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(g.key),
	})

	return &keyPem, nil
}

func (g *TLSCertGenerator) GenerateDefaultPEMCertificate() (*[]byte, error) {

	sn, err := g.randomBigInt()
	if err != nil {
		return nil, err
	}
	tml := x509.Certificate{
		// you can add any attr that you need
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(5, 0, 0),
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName:   x509CertificateCommonName,
			Organization: []string{x509CertificateOrganization},
		},
		BasicConstraintsValid: true,
	}

	return g.GeneratePEMCertificate(tml)
}

func (g *TLSCertGenerator) randomBigInt() (*big.Int, error) {
	//Max random value, a 130-bits integer, i.e 2^130 - 1
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))

	//Generate cryptographically strong pseudo-random between 0 - max
	n, err := rand.Int(rand.Reader, max)

	return n, err
}

func (g *TLSCertGenerator) GeneratePEMCertificate(x509Cert x509.Certificate) (*[]byte, error){

	cert, err := x509.CreateCertificate(rand.Reader, &x509Cert, &x509Cert, &g.key.PublicKey, g.key)
	if err != nil {
		return nil, fmt.Errorf("certificate cannot be created: %v", err.Error())
	}

	// Generate a pem block with the certificate
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})

	return &certPem, nil
}

func (g *TLSCertGenerator) GenerateX509KeyPair(certPEM, keyPEM []byte) (*tls.Certificate, error) {

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, fmt.Errorf("cannot be loaded the certificate: %v", err.Error())
	}

	return &tlsCert, nil
}

func (g *TLSCertGenerator) GenerateDefaultX509KeyPair() (*tls.Certificate, error) {

	keyPEM, err := g.GeneratePEMPrivateKey()
	if err != nil {
		return nil, err
	}

	certPEM, err := g.GenerateDefaultPEMCertificate()
	if err != nil {
		return nil, err
	}

	return g.GenerateX509KeyPair(*certPEM, *keyPEM)
}