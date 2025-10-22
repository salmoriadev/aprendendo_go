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

	estrategiaChave := criptografia.NovaEstrategiaChaveRSA()
	estrategiaCertificado := criptografia.NovaEstrategiaCertificado()
	estrategiaResumo := criptografia.NovaEstrategiaResumoSha256()
	estrategiaAssinatura := criptografia.NovaEstrategiaAssinaturaPkcs1v15()

	chaveAc := servicos.ExecucaoChaves(tamanhoChaves,
		caminho, estrategiaChave)
	chaveUsuario := servicos.ExecucaoChaves(tamanhoChaves,
		caminho, estrategiaChave)
	servicos.ExecucaoCertificados(chaveAc, chaveUsuario,
		validadeCertAC, validadeCert, caminho, "UFSC",
		"BR", "Santa Catarina", "Florianopolis",
		"localhost", estrategiaCertificado)

	caminhoArquivoTxt := caminho + "/mensagem.txt"
	conteudoArquivo := "Esta é uma mensagem importante."

	servicos.GerarArquivoPDF(
		caminhoArquivoTxt, caminho+"/mensagem.pdf",
		conteudoArquivo)

	servicos.ResumirPDF(caminho+"/mensagem.pdf", caminho+"/resumo.txt",
		estrategiaResumo)

	assinatura := servicos.AssinarDocumentoPDF(caminho+"/mensagem.pdf",
		caminho+"/assinatura.txt", &chaveUsuario,
		estrategiaResumo, estrategiaAssinatura)

	fmt.Println("Assinatura do documento PDF gerada com sucesso.")

	ehValida := servicos.VerificarAssinaturaDocumentoPDF(
		caminho+"/mensagem.pdf", assinatura,
		chaveUsuario.ChavePublica,
		estrategiaResumo, estrategiaAssinatura)

	if ehValida {
		fmt.Println("A assinatura é válida.")
	} else {
		fmt.Println("A assinatura não é válida.")
	}
}
