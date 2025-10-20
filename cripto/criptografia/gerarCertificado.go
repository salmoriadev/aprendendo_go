package criptografia

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

type Certificado struct {
	CertificadoBytes []byte
	Certificado      *x509.Certificate
}

func NovoCertificado() Certificado {
	return Certificado{}
}

func CertificadoParaPEM(certificado *Certificado) []byte {
	blocoPEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificado.CertificadoBytes,
	}
	return pem.EncodeToMemory(blocoPEM)
}

func GerarCertificadoAutoassinado(chavePrivada *rsa.PrivateKey,
	sujeito *pkix.Name, validadeEmAnos int) (Certificado, error) {
	var cert Certificado = NovoCertificado()
	inicioPrazo := time.Now()
	validade := inicioPrazo.AddDate(validadeEmAnos, 0, 0)
	var permissoesDaChave x509.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	var propositosDaChave []x509.ExtKeyUsage = []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}

	cert.Certificado = &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               *sujeito,
		NotBefore:             inicioPrazo,
		NotAfter:              validade,
		KeyUsage:              permissoesDaChave,
		ExtKeyUsage:           propositosDaChave,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	var err error
	cert.CertificadoBytes, err = x509.CreateCertificate(rand.Reader,
		cert.Certificado, cert.Certificado,
		&chavePrivada.PublicKey, chavePrivada)
	if err != nil {
		return cert, err
	}

	return cert, nil
}

func GerarCertificadoAssinadoPorAC(chavePrivadaSujeito *rsa.PrivateKey,
	sujeito pkix.Name, validadeEmAnos int, certPai *x509.Certificate,
	chavePrivadaPai *rsa.PrivateKey) (Certificado, error) {

	var cert Certificado = NovoCertificado()
	inicioPrazo := time.Now()
	validade := inicioPrazo.AddDate(validadeEmAnos, 0, 0)
	var permissoesDaChave x509.KeyUsage = x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature
	var propositosDaChave []x509.ExtKeyUsage = []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth}

	cert.Certificado = &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      sujeito,
		NotBefore:    inicioPrazo,
		NotAfter:     validade,
		KeyUsage:     permissoesDaChave,
		ExtKeyUsage:  propositosDaChave,
		IsCA:         false,
	}

	var err error
	cert.CertificadoBytes, err = x509.CreateCertificate(
		rand.Reader, cert.Certificado, certPai,
		&chavePrivadaSujeito.PublicKey, chavePrivadaPai)
	if err != nil {
		return cert, err
	}
	return cert, nil
}
