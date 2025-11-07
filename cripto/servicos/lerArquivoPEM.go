/*
Utilitários voltados à leitura de chaves e certificados em formato PEM,
permitindo reutilização de artefatos gerados pelo fluxo em outros cenários.
*/
package servicos

import (
	"cripto/criptografia"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

// LerChavesDeArquivoPEM percorre um arquivo PEM e carrega chaves pública/privada em memória.
func LerChavesDeArquivoPEM(caminho string) (criptografia.ParDeChaves, error) {
	chaves := criptografia.ParDeChaves{}
	arquivo, err := os.Open(caminho)
	if err != nil {
		return chaves, err
	}
	defer func(arquivo *os.File) {
		err := arquivo.Close()
		if err != nil {
			fmt.Println("Erro ao fechar arquivo: ", err)
		}
	}(arquivo)

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

// LerCertificadoDeArquivoPEM abre um arquivo PEM e reconstrói a estrutura de
// certificado X.509 correspondente.
func LerCertificadoDeArquivoPEM(caminho string) (
	criptografia.Certificado, error) {
	cert := criptografia.Certificado{}
	data, err := os.ReadFile(caminho)
	if err != nil {
		return cert, err
	}

	bloco, _ := pem.Decode(data)
	if bloco == nil || bloco.Type != "CERTIFICATE" {
		return cert, fmt.Errorf(
			"arquivo PEM inválido ou sem certificado")
	}

	cert.CertificadoBytes = bloco.Bytes
	cert.Certificado, err = x509.ParseCertificate(cert.CertificadoBytes)
	if err != nil {
		return cert, err
	}

	return cert, nil
}
