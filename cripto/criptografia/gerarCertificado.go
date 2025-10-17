package criptografia

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

type certificado struct {
	CertificadoBytes []byte
	Certificado      *x509.Certificate
}

func NovoCertificado() certificado {
	return certificado{}
}

func GerarCertificadoAutoassinado(chavePrivada *rsa.PrivateKey, sujeito pkix.Name,
	validadeEmAnos int) (certificado, error) {
	var cert certificado = NovoCertificado()
	inicioPrazo := time.Now()
	validade := inicioPrazo.AddDate(validadeEmAnos, 0, 0)
	var permissoesDaChave x509.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	var propositosDaChave []x509.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth,
		x509.ExtKeyUsageClientAuth}

	cert.Certificado = &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               sujeito,
		NotBefore:             inicioPrazo,
		NotAfter:              validade,
		KeyUsage:              permissoesDaChave,
		ExtKeyUsage:           propositosDaChave,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	var err error
	cert.CertificadoBytes, err = x509.CreateCertificate(rand.Reader, cert.Certificado,
		cert.Certificado, &chavePrivada.PublicKey, chavePrivada)
	if err != nil {
		return cert, err
	}

	return cert, nil
}

func GerarCertificadoAssinadoPorAC(chavePrivadaSujeito *rsa.PrivateKey, sujeito pkix.Name,
	validadeEmAnos int, certPai *x509.Certificate,
	chavePrivadaPai *rsa.PrivateKey) (certificado, error) {

	var cert certificado = NovoCertificado()
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

func EscreverCertificadoParaArquivoPEM(certificado certificado, caminhoCompleto string) error {
	diretorio := filepath.Dir(caminhoCompleto)

	if err := os.MkdirAll(diretorio, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório %s: %w", diretorio, err)
	}

	// Cria o diretório (e todos os pais necessários) se ele não existir
	// 0755 são as permissões padrão para um diretório
	if err := os.MkdirAll(diretorio, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório %s: %w", diretorio, err)
	}

	// --- O RESTO DO SEU CÓDIGO CONTINUA IGUAL ---

	blocoPEM := &pem.Block{Type: "CERTIFICATE", Bytes: certificado.CertificadoBytes}
	dadosPEM := pem.EncodeToMemory(blocoPEM)

	err := os.WriteFile(caminhoCompleto, dadosPEM, 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever arquivo PEM: %w", err)
	}

	return nil
}
