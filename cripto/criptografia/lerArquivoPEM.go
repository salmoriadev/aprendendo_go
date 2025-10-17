package criptografia

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

func LerChavesDeArquivoPEM(caminho string) (parDeChaves, error) {
	chaves := NovaParDeChaves()
	arquivo, err := os.Open(caminho)
	if err != nil {
		return chaves, err
	}
	defer arquivo.Close()

	contents, err := io.ReadAll(arquivo)
	if err != nil {
		return chaves, err
	}

	for {
		var block *pem.Block
		block, contents = pem.Decode(contents)
		if block == nil {
			break
		}
		switch block.Type {
		case "PRIVATE KEY":
			chaves.ChavePrivada, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		case "PUBLIC KEY":
			chaves.ChavePublica, err = x509.ParsePKCS1PublicKey(block.Bytes)
		}
		if err != nil {
			return chaves, err
		}
	}
	return chaves, nil
}

func LerCertificadoDeArquivoPEM(caminho string) (certificado, error) {
	cert := NovoCertificado()
	data, err := os.ReadFile(caminho)
	if err != nil {
		return cert, err
	}

	bloco, _ := pem.Decode(data)
	if bloco == nil || bloco.Type != "CERTIFICATE" {
		return cert, fmt.Errorf("arquivo PEM inválido ou sem certificado")
	}

	cert.CertificadoBytes = bloco.Bytes
	cert.Certificado, err = x509.ParseCertificate(cert.CertificadoBytes)
	if err != nil {
		return cert, err
	}

	return cert, nil
}
