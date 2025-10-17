package main

import (
	"cripto/criptografia"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
)

func execucaoChaves() ([]byte, []byte) {
	chaves, err := criptografia.GerarChavePrivada(2048)
	if err != nil {
		fmt.Println("Erro ao gerar chave privada:", err)
		return nil, nil
	}
	bytesChavePrivada := x509.MarshalPKCS1PrivateKey(chaves.ChavePrivada)
	bytesChavePublica := x509.MarshalPKCS1PublicKey(chaves.ChavePublica)

	fmt.Println("Chave privada gerada com sucesso:")
	fmt.Println(criptografia.CodificarChaveParaBase64(bytesChavePrivada))

	fmt.Println("Chave pública gerada com sucesso:")
	fmt.Println(criptografia.CodificarChaveParaBase64(bytesChavePublica))
	return bytesChavePrivada, bytesChavePublica
}

func execucaoCertificados() {
	fmt.Println("--- Etapa 1: Gerando a AC Raiz (Root CA) ---")

	chaveAC, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Erro ao gerar chave para AC: %v", err)
	}

	sujeitoAC := pkix.Name{
		Organization: []string{"Minha Empresa AC"},
		Country:      []string{"BR"},
		Province:     []string{"Santa Catarina"},
		Locality:     []string{"Florianopolis"},
		CommonName:   "Minha AC Raiz",
	}

	certAC, err := criptografia.GerarCertificadoAutoassinado(chaveAC, sujeitoAC, 10)
	if err != nil {
		log.Fatalf("Erro ao gerar certificado autoassinado para AC: %v", err)
	}

	fmt.Println("Certificado da AC Raiz gerado com sucesso!")

	pemAC := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certAC.CertificadoBytes})
	fmt.Println(string(pemAC))

	fmt.Println("\n--- Etapa 2: Gerando Certificado de Servidor ---")

	chaveServidor, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Erro ao gerar chave para o Servidor: %v", err)
	}

	sujeitoServidor := pkix.Name{
		Organization: []string{"Meu Site"},
		Country:      []string{"BR"},
		Province:     []string{"Santa Catarina"},
		Locality:     []string{"Florianopolis"},
		CommonName:   "localhost",
	}

	certServidor, err := criptografia.GerarCertificadoAssinadoPorAC(chaveServidor, sujeitoServidor, 1, certAC.Certificado, chaveAC)
	if err != nil {
		log.Fatalf("Erro ao gerar certificado de servidor: %v", err)
	}

	fmt.Println("Certificado do Servidor gerado e assinado com sucesso!")

	pemServidor := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certServidor.CertificadoBytes})
	fmt.Println(string(pemServidor))
}

func main() {
	execucaoChaves()
	execucaoCertificados()
}
