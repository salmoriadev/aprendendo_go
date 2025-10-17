package criptografia

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
)

func ExecucaoChaves(tamanhoChave int, caminho string) parDeChaves {
	chaves, err := GerarChavePrivada(tamanhoChave)
	if err != nil {
		fmt.Println("Erro ao gerar chave privada:", err)
		return parDeChaves{}
	}
	bytesChavePrivada := x509.MarshalPKCS1PrivateKey(chaves.ChavePrivada)
	bytesChavePublica := x509.MarshalPKCS1PublicKey(chaves.ChavePublica)

	fmt.Println("Chave privada gerada com sucesso:")
	fmt.Println(CodificarChaveParaBase64(bytesChavePrivada))

	fmt.Println("Chave pública gerada com sucesso:")
	fmt.Println(CodificarChaveParaBase64(bytesChavePublica))

	EscreverChavePrivadaParaArquivoPEM(
		chaves.ChavePrivada, caminho+"/chave_privada.pem")
	EscreverChavePublicaParaArquivoPEM(
		chaves.ChavePublica, caminho+"/chave_publica.pem")

	return chaves
}

func ExecucaoCertificados(chaveAC parDeChaves, chaveCert parDeChaves,
	tamanhoChave int, validadeCertAC int,
	validadeCert int, caminho string,
	organizacao string, pais string,
	provincia string, localidade string,
	nomeComum string) {
	fmt.Println("--- Etapa 1: Gerando a AC Raiz (Root CA) ---")

	sujeitoAC := pkix.Name{
		Organization: []string{"AC"},
		Country:      []string{"BR"},
		Province:     []string{"Santa Catarina"},
		Locality:     []string{"Florianopolis"},
		CommonName:   "Minha AC Raiz",
	}
	certAC, err := GerarCertificadoAutoassinado(
		chaveAC.ChavePrivada, sujeitoAC, validadeCertAC)
	if err != nil {
		log.Fatalf("Erro ao gerar certificado autoassinado para AC: %v", err)
	}

	fmt.Println("Certificado da AC Raiz gerado")
	err = EscreverCertificadoParaArquivoPEM(
		certAC, caminho+"/certificado_ac.pem")
	if err != nil {
		log.Fatalf(
			"Erro ao escrever certificado da AC Raiz em arquivo: %v", err)
	}
	fmt.Println("Certificado da AC Raiz escrito em arquivo com sucesso!")

	fmt.Println("\n--- Etapa 2: Gerando Certificado de Servidor ---")

	if err != nil {
		log.Fatalf("Erro ao gerar chave para o Servidor: %v", err)
	}

	sujeitoServidor := pkix.Name{
		Organization: []string{organizacao},
		Country:      []string{pais},
		Province:     []string{provincia},
		Locality:     []string{localidade},
		CommonName:   nomeComum,
	}

	certServidor, err := GerarCertificadoAssinadoPorAC(
		chaveCert.ChavePrivada, sujeitoServidor,
		validadeCert, certAC.Certificado, chaveAC.ChavePrivada)
	if err != nil {
		log.Fatalf("Erro ao gerar certificado de servidor: %v", err)
	}

	err = EscreverCertificadoParaArquivoPEM(
		certServidor, caminho+"/certificado_servidor.pem")
	if err != nil {
		log.Fatalf("Erro ao escrever certificado do Servidor em arquivo: %v", err)
	}
	fmt.Println("Certificado do Servidor gerado e assinado com sucesso!")
	pemBloco := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certServidor.CertificadoBytes,
	}
	pemServidor := pem.EncodeToMemory(pemBloco)
	fmt.Println(string(pemServidor))
}
