package main

import (
	"cripto/criptografia"
	"cripto/servicos"
	"fmt"
)

func main() {
	tamanhoChaves := 2048
	validadeCertAC := 10
	validadeCert := 1
	chaveAc := servicos.ExecucaoChaves(tamanhoChaves, caminho)
	chaveUsuario := servicos.ExecucaoChaves(tamanhoChaves, caminho)
	servicos.ExecucaoCertificados(chaveAc, chaveUsuario, tamanhoChaves,
		validadeCertAC, validadeCert, caminho, "UFSC",
		"BR", "Santa Catarina", "Florianopolis", "localhost")

	caminhoArquivoTxt := caminho + "/mensagem.txt"
	conteudoArquivo := "Esta é uma mensagem importante."

	servicos.GerarArquivoPDF(caminhoArquivoTxt, caminho+"/mensagem.pdf", conteudoArquivo)

	var estrategiaResumo criptografia.EstrategiaResumo = &criptografia.ResumoSha256{}
	var estrategiaAssinatura criptografia.EstrategiaAssinatura = &criptografia.AssinaturaPkcs1v15{}

	servicos.ResumirPDF(caminho+"/mensagem.pdf", caminho+"/resumo.txt",
		estrategiaResumo)

	assinatura := servicos.AssinarDocumentoPDF(caminho+"/mensagem.pdf", caminho+"/assinatura.txt", chaveUsuario,
		estrategiaResumo, estrategiaAssinatura)

	fmt.Println("Assinatura do documento PDF gerada com sucesso.")

	valida := servicos.VerificarAssinaturaDocumentoPDF(caminho+"/mensagem.pdf", assinatura, chaveUsuario.ChavePublica,
		estrategiaResumo, estrategiaAssinatura)

	if valida {
		fmt.Println("A assinatura é válida.")
	} else {
		fmt.Println("A assinatura não é válida.")
	}
}
