/*
Aplicacao principal responsavel por coordenar o fluxo de certificação digital.
Aqui definimos as estratégias utilizadas (Strategy Pattern), orquestramos a
emissão de certificados, geração do PDF, assinatura e verificação mostrando o
valor do portfólio para o LabSEC.
*/
package main

import (
	"cripto/criptografia"
	"cripto/servicos"
	"fmt"
	"log"
	"path/filepath"
)

// AplicacaoAssinaturaDigital demonstra um fluxo completo de ICP com estratégias plugáveis.
// Os comentários explicam como cada camada se integra para facilitar a avaliação do portfólio.
// ConfiguracaoCriptografia evidencia o uso do Strategy Pattern ao
// centralizar as escolhas de algoritmo.
type ConfiguracaoCriptografia struct {
	TamanhoChaveBits          int
	ValidadeCertificadoACAnos int
	ValidadeCertificadoAnos   int
	EstrategiaChave           criptografia.IEstrategiaChave
	EstrategiaCertificado     criptografia.IEstrategiaCertificado
	EstrategiaResumo          criptografia.IEstrategiaResumo
	EstrategiaAssinatura      criptografia.IEstrategiaAssinatura
}

// novaConfiguracaoPadrao retorna as escolhas padrão usadas no fluxo
// quando nenhuma configuração é fornecida.
func novaConfiguracaoPadrao() ConfiguracaoCriptografia {
	return ConfiguracaoCriptografia{
		TamanhoChaveBits:          2048,
		ValidadeCertificadoACAnos: 10,
		ValidadeCertificadoAnos:   1,
		EstrategiaChave:           criptografia.NovaEstrategiaChaveRSA(),
		EstrategiaCertificado:     criptografia.NovaEstrategiaCertificado(),
		EstrategiaResumo:          criptografia.NovaEstrategiaResumoSha256(),
		EstrategiaAssinatura:      criptografia.NovaEstrategiaAssinaturaPkcs1v15(),
	}
}

// main executa o fluxo completo: gera chaves e certificados, prepara o PDF,
// assina e valida o resultado.
func main() {
	config := novaConfiguracaoPadrao()

	chaveAutoridade, err := servicos.GerarParDeChaves(config.TamanhoChaveBits,
		caminho, "autoridade", config.EstrategiaChave)
	if err != nil {
		log.Fatalf("erro ao preparar chaves da autoridade: %v", err)
	}

	chaveUsuario, err := servicos.GerarParDeChaves(config.TamanhoChaveBits,
		caminho, "usuario", config.EstrategiaChave)
	if err != nil {
		log.Fatalf("erro ao preparar chaves do usuário: %v", err)
	}

	dadosAutoridade := servicos.DadosIdentificacaoCertificado{
		Organizacao: "UFSC",
		Pais:        "BR",
		Provincia:   "Santa Catarina",
		Localidade:  "Florianopolis",
		NomeComum:   "UFSC Autoridade Certificadora",
	}

	dadosUsuario := servicos.DadosIdentificacaoCertificado{
		Organizacao: "UFSC",
		Pais:        "BR",
		Provincia:   "Santa Catarina",
		Localidade:  "Florianopolis",
		NomeComum:   "localhost",
	}

	if err := servicos.GerarCertificados(chaveAutoridade, chaveUsuario,
		config.ValidadeCertificadoACAnos, config.ValidadeCertificadoAnos,
		caminho, dadosAutoridade,
		dadosUsuario, config.EstrategiaCertificado); err != nil {
		log.Fatalf("erro ao gerar certificados: %v", err)
	}

	caminhoArquivoTxt := filepath.Join(caminho, "mensagem.txt")
	conteudoArquivo := "Esta é uma mensagem importante."

	caminhoArquivoPdf := filepath.Join(caminho, "mensagem.pdf")
	caminhoResumo := filepath.Join(caminho, "resumo.txt")
	caminhoAssinatura := filepath.Join(caminho, "assinatura.txt")

	if err := servicos.GerarArquivoPDF(
		caminhoArquivoTxt, caminhoArquivoPdf,
		conteudoArquivo); err != nil {
		log.Fatalf("erro ao preparar PDF: %v", err)
	}

	if _, err := servicos.ResumirPDF(caminhoArquivoPdf, caminhoResumo,
		config.EstrategiaResumo); err != nil {
		log.Fatalf("erro ao resumir PDF: %v", err)
	}

	assinaturaDigital, err := servicos.AssinarDocumentoPDF(caminhoArquivoPdf,
		caminhoAssinatura, &chaveUsuario,
		config.EstrategiaResumo, config.EstrategiaAssinatura)
	if err != nil {
		log.Fatalf("erro ao assinar PDF: %v", err)
	}

	fmt.Println("Assinatura digital do documento PDF gerada com sucesso.")

	assinaturaValida, err := servicos.VerificarAssinaturaDocumentoPDF(
		caminhoArquivoPdf, assinaturaDigital,
		chaveUsuario.ChavePublica,
		config.EstrategiaResumo, config.EstrategiaAssinatura)
	if err != nil {
		log.Fatalf("erro ao verificar assinatura: %v", err)
	}

	if assinaturaValida {
		fmt.Println("A assinatura digital é válida.")
	} else {
		fmt.Println("A assinatura digital não é válida.")
	}
}
