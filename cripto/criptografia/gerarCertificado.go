package criptografia

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

type certificado struct {
	CertificadoBytes []byte
	Certificado      *x509.Certificate
}

func NovoCertificado() certificado {
	return certificado{}
}

func GerarCertificadoAutoassinado(chavePrivada *rsa.PrivateKey, sujeito pkix.Name, validadeEmAnos int) (certificado, error) {
	var cert certificado = NovoCertificado()
	inicioPrazo := time.Now()
	validade := inicioPrazo.AddDate(validadeEmAnos, 0, 0)

	cert.Certificado = &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               sujeito,
		NotBefore:             inicioPrazo,
		NotAfter:              validade,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	var err error
	cert.CertificadoBytes, err = x509.CreateCertificate(rand.Reader, cert.Certificado, cert.Certificado, &chavePrivada.PublicKey, chavePrivada)
	if err != nil {
		return cert, err
	}

	return cert, nil
}

func GerarCertificadoAssinadoPorAC(chavePrivadaSujeito *rsa.PrivateKey, sujeito pkix.Name, validadeEmAnos int, certPai *x509.Certificate, chavePrivadaPai *rsa.PrivateKey) (certificado, error) {
	var cert certificado = NovoCertificado()
	inicioPrazo := time.Now()
	validade := inicioPrazo.AddDate(validadeEmAnos, 0, 0)

	cert.Certificado = &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      sujeito,
		NotBefore:    inicioPrazo,
		NotAfter:     validade,
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IsCA:         false,
	}

	var err error
	cert.CertificadoBytes, err = x509.CreateCertificate(rand.Reader, cert.Certificado, certPai, &chavePrivadaSujeito.PublicKey, chavePrivadaPai)
	if err != nil {
		return cert, err
	}
	return cert, nil
}

func EscreverCertificadoParaArquivoPEM(cert certificado, caminho string) error {
	arquivo, err := os.Create(caminho)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	err = pem.Encode(arquivo, &pem.Block{Type: "CERTIFICATE", Bytes: cert.CertificadoBytes})
	if err != nil {
		return err
	}

	return nil
}
