package criptografia

import (
	"crypto/rsa"
	"crypto/x509/pkix"
	"encoding/hex"
	"fmt"
	"log"
)

func ExecucaoChaves(tamanhoChave int, caminho string) parDeChaves {
	chaves, err := GerarChavePrivada(tamanhoChave)
	if err != nil {
		fmt.Println("Erro ao gerar chave privada:", err)
		return parDeChaves{}
	}

	fmt.Println("Chave privada e pública geradas com sucesso:")

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

	fmt.Println("\n--- Etapa 2: Gerando Certificado do usuário ---")

	if err != nil {
		log.Fatalf("Erro ao gerar chave para o usuário: %v", err)
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
		log.Fatalf("Erro ao gerar certificado de usuário: %v", err)
	}

	err = EscreverCertificadoParaArquivoPEM(
		certServidor, caminho+"/certificado_usuario.pem")
	if err != nil {
		log.Fatalf("Erro ao escrever certificado do usuário em arquivo: %v", err)
	}
	fmt.Println("Certificado do usuário gerado e assinado com sucesso!")
}
func GerarArquivoPDF(caminhoArquivoTxt string, caminhoArquivoPdf string, conteudoArquivo string) {
	GerarArquivoTexto(caminhoArquivoTxt, conteudoArquivo)
	TxtParaPdf(caminhoArquivoTxt, caminhoArquivoPdf)
}
func ResumirPDF(caminhoArquivoPdf string, caminhoResumo string) {
	arrayBytes, err := PdfParaBytes(caminhoArquivoPdf)
	if err != nil {
		log.Fatalf("Erro ao ler PDF: %v", err)
	}
	resumo := hex.EncodeToString(GerarResumoCriptografico(arrayBytes))
	fmt.Println("Resumo criptográfico do PDF:", resumo)
	GerarArquivoTexto(caminhoResumo, resumo)
}

func AssinarDocumentoPDF(caminhoArquivoPdf string, caminhoAssinatura string, chaves parDeChaves) []byte {
	arrayBytes, err := PdfParaBytes(caminhoArquivoPdf)
	if err != nil {
		log.Fatalf("Erro ao ler PDF: %v", err)
	}

	resumoEmBytes := GerarResumoCriptografico(arrayBytes)

	assinatura, err := AssinarResumo(resumoEmBytes, chaves.ChavePrivada)
	if err != nil {
		log.Fatalf("Erro ao assinar o resumo: %v", err)
	}
	GerarArquivoTexto(caminhoAssinatura, hex.EncodeToString(assinatura))
	return assinatura

}

func VerificarAssinaturaDocumentoPDF(caminhoArquivoPdf string, assinatura []byte, chavePublica *rsa.PublicKey) bool {
	arrayBytes, err := PdfParaBytes(caminhoArquivoPdf)
	if err != nil {
		log.Fatalf("Erro ao ler PDF: %v", err)
	}
	resumo := GerarResumoCriptografico(arrayBytes)
	resumoEmBytes := []byte(resumo)
	err = VerificarAssinatura(resumoEmBytes, assinatura, chavePublica)
	if err != nil {
		log.Fatalf("Erro ao verificar a assinatura: %v", err)
		return false
	}
	return true
}
