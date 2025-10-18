package main

import (
	"cripto/criptografia"
	"fmt"
)

func main() {
	tamanhoChaves := 2048
	validadeCertAC := 10
	validadeCert := 1

	chaveAc := criptografia.ExecucaoChaves(tamanhoChaves, caminho)
	chaveUsuario := criptografia.ExecucaoChaves(tamanhoChaves, caminho)
	criptografia.ExecucaoCertificados(chaveAc, chaveUsuario, tamanhoChaves,
		validadeCertAC, validadeCert, caminho, "Meu Site",
		"BR", "Santa Catarina", "Florianopolis", "localhost")
	caminhoArquivoTxt := caminho + "/mensagem.txt"
	conteudoArquivo := "Esta é uma mensagem importante."
	criptografia.GerarArquivoPDF(caminhoArquivoTxt, caminho+"/mensagem.pdf", conteudoArquivo)
	criptografia.ResumirPDF(caminho+"/mensagem.pdf", caminho+"/resumo.txt")
	assinatura := criptografia.AssinarDocumentoPDF(caminho+"/mensagem.pdf", caminho+"/assinatura.txt", chaveUsuario)
	fmt.Println("Assinatura do documento PDF gerada com sucesso.")
	valida := criptografia.VerificarAssinaturaDocumentoPDF(caminho+"/mensagem.pdf", assinatura, chaveUsuario.ChavePublica)
	if valida {
		fmt.Println("A assinatura é válida.")
	} else {
		fmt.Println("A assinatura não é válida.")
	}
}
