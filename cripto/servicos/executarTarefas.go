package servicos

import (
	"cripto/criptografia"
	"crypto/x509/pkix"
	"fmt"
	"log"
)

func ExecucaoChaves(tamanhoChave int, caminho string,
	estrategia criptografia.EstrategiaChave) criptografia.ParDeChaves {

	chaves, err := estrategia.GerarChavePrivada(tamanhoChave)
	if err != nil {
		fmt.Println("Erro ao gerar chave privada: ", err)
		return criptografia.ParDeChaves{}
	}

	fmt.Println("Chave privada e pública geradas com sucesso")

	dadosPEMPriv := criptografia.ChavePrivadaParaPEM(chaves.ChavePrivada)
	err = escreverArquivo(caminho+"/chave_privada.pem", dadosPEMPriv)
	if err != nil {
		log.Fatalf("Erro ao escrever chave privada: %v", err)
	}

	dadosPEMPub := criptografia.ChavePublicaParaPEM(chaves.ChavePublica)
	err = escreverArquivo(caminho+"/chave_publica.pem", dadosPEMPub)
	if err != nil {
		log.Fatalf("Erro ao escrever chave publica: %v", err)
	}

	return chaves
}

func ExecucaoCertificados(chaveAC, chaveCert criptografia.ParDeChaves,
	tamanhoChave int, validadeCertAC int, validadeCert int, caminho string,
	organizacao, pais, provincia, localidade, nomeComum string,
	estrategia criptografia.EstrategiaCertificado) {

	sujeitoAC := pkix.Name{
		Organization: []string{"UFC"},
		Country:      []string{"BR"},
		Province:     []string{"Sao Paulo"},
		Locality:     []string{"Sao Bernardo do Campo"},
		CommonName:   "chama",
	}

	certAC, err := estrategia.GerarCertificadoAutoassinado(
		chaveAC.ChavePrivada, &sujeitoAC, validadeCertAC)
	if err != nil {
		log.Fatalf("Erro ao gerar certificado autoassinado para AC: %v", err)
	}

	dadosPEMCertAC := criptografia.CertificadoParaPEM(&certAC)
	err = escreverArquivo(caminho+"/certificado_ac.pem", dadosPEMCertAC)
	if err != nil {
		log.Fatalf("Erro ao escrever certificado da AC Raiz: %v", err)
	}
	fmt.Println("Certificado da AC Raiz escrito em arquivo com sucesso!")

	sujeitoUsuario := pkix.Name{
		Organization: []string{organizacao},
		Country:      []string{pais},
		Province:     []string{provincia},
		Locality:     []string{localidade},
		CommonName:   nomeComum,
	}

	certUsuario, err := estrategia.GerarCertificadoAssinadoPorAC(
		chaveCert.ChavePrivada, sujeitoUsuario,
		validadeCert, certAC.Certificado, chaveAC.ChavePrivada)
	if err != nil {
		log.Fatalf("Erro ao gerar certificado de usuário: %v", err)
	}

	dadosPEMCertUsuario := criptografia.CertificadoParaPEM(&certUsuario)
	err = escreverArquivo(caminho+"/certificado_usuario.pem", dadosPEMCertUsuario)
	if err != nil {
		log.Fatalf("Erro ao escrever certificado do usuário: %v", err)
	}
	fmt.Println("Certificado do usuário gerado e assinado com sucesso!")
}
