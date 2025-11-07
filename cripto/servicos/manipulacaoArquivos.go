/*
Serviços auxiliares de manipulação de arquivos que dão suporte ao fluxo
criptográfico: criação de TXT/PDF, geração de hashes, assinatura e verificação
de documentos persistidos.
*/
package servicos

import (
	"cripto/criptografia"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

// Package servicos fornece funcoes de manipulacao de arquivos para suportar o fluxo criptografico.
// Mantem entradas e saidas separadas das estrategias puras definidas no pacote criptografia.

// GerarArquivoPDF cria um TXT com o conteudo fornecido e o converte em PDF.
// Ideal para preparar um documento de exemplo antes de aplicar as etapas criptograficas.
func GerarArquivoPDF(caminhoArquivoTxt string, caminhoArquivoPdf string,
	conteudoArquivo string) error {
	if err := gerarArquivoTexto(caminhoArquivoTxt, conteudoArquivo); err != nil {
		return fmt.Errorf("erro ao gerar arquivo TXT: %w", err)
	}
	if err := txtParaPdf(caminhoArquivoTxt, caminhoArquivoPdf); err != nil {
		return fmt.Errorf("erro ao converter TXT para PDF: %w", err)
	}
	return nil
}

// ResumirPDF gera o hash hexadecimal do PDF e o salva, devolvendo o resumo para uso posterior.
func ResumirPDF(caminhoArquivoPdf string, caminhoResumo string,
	estrategiaResumo criptografia.IEstrategiaResumo) (string, error) {
	dados, err := pdfParaBytes(caminhoArquivoPdf)
	if err != nil {
		return "", fmt.Errorf("erro ao ler PDF: %w", err)
	}
	resumo := hex.EncodeToString(estrategiaResumo.Resumir(dados))

	fmt.Println("Resumo criptográfico do PDF:", resumo)
	if err = gerarArquivoTexto(caminhoResumo, resumo); err != nil {
		return "", fmt.Errorf("erro ao salvar resumo: %w", err)
	}
	return resumo, nil
}

// AssinarDocumentoPDF assina o resumo do PDF em memoria e guarda o resultado em disco.
// Retorna a assinatura como slice de bytes, facilitando a verificacao sequencial.
func AssinarDocumentoPDF(caminhoArquivoPdf string, caminhoAssinatura string,
	chaves *criptografia.ParDeChaves,
	estrategiaResumo criptografia.IEstrategiaResumo,
	estrategiaAssinatura criptografia.IEstrategiaAssinatura) ([]byte, error) {

	dados, err := pdfParaBytes(caminhoArquivoPdf)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler PDF: %w", err)
	}

	resumoEmBytes := estrategiaResumo.Resumir(dados)
	hashFunc := estrategiaResumo.HashFunc()

	assinatura, err := estrategiaAssinatura.Assinar(resumoEmBytes,
		chaves.ChavePrivada, hashFunc)
	if err != nil {
		return nil, fmt.Errorf("erro ao assinar resumo criptográfico: %w", err)
	}

	if err = gerarArquivoTexto(caminhoAssinatura,
		hex.EncodeToString(assinatura)); err != nil {
		return nil, fmt.Errorf("erro ao persistir assinatura: %w", err)
	}
	return assinatura, nil
}

// VerificarAssinaturaDocumentoPDF confirma se a assinatura confere com o resumo calculado.
// Informa o resultado booleano mantendo a mensagem de erro detalhada para o chamador.
func VerificarAssinaturaDocumentoPDF(caminhoArquivoPdf string,
	assinatura []byte, chavePublica *rsa.PublicKey,
	estrategiaResumo criptografia.IEstrategiaResumo,
	estrategiaAssinatura criptografia.IEstrategiaAssinatura) (bool, error) {

	dados, err := pdfParaBytes(caminhoArquivoPdf)
	if err != nil {
		return false, fmt.Errorf("erro ao ler PDF: %w", err)
	}

	resumoEmBytes := estrategiaResumo.Resumir(dados)
	hashFunc := estrategiaResumo.HashFunc()

	if err = estrategiaAssinatura.VerificarAssinatura(resumoEmBytes,
		assinatura, chavePublica, hashFunc); err != nil {
		return false, fmt.Errorf("erro ao verificar assinatura digital: %w", err)
	}
	return true, nil
}

// gerarArquivoTexto garante a existencia do diretorio e, em seguida, grava o conteudo puro.
func gerarArquivoTexto(caminho string, conteudo string) error {
	diretorio := filepath.Dir(caminho)
	if err := os.MkdirAll(diretorio, 0755); err != nil {
		return fmt.Errorf(
			"erro ao criar diretório %s: %w", diretorio, err)
	}
	return os.WriteFile(caminho, []byte(conteudo), 0644)
}

// pdfParaBytes le o PDF diretamente em memoria, preservando o conteudo para hash e assinatura.
func pdfParaBytes(caminhoArquivoPdf string) ([]byte, error) {
	return os.ReadFile(caminhoArquivoPdf)
}

// txtParaPdf converte o TXT para PDF simples, mantendo a fonte padrao do gofpdf.
func txtParaPdf(caminhoTxt string, caminhoPdf string) error {
	text, err := os.ReadFile(caminhoTxt)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo TXT: %w", err)
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 14)
	width, _ := pdf.GetPageSize()
	pdf.MultiCell(width-20, 10, string(text), "", "L", false)
	if err = pdf.OutputFileAndClose(caminhoPdf); err != nil {
		return fmt.Errorf("erro ao gerar PDF: %w", err)
	}
	fmt.Println("PDF gerado com sucesso")
	return nil
}

// escreverArquivo encapsula a escrita de bytes garantindo o diretorio alvo.
func escreverArquivo(caminho string, dados []byte) error {
	diretorio := filepath.Dir(caminho)
	if err := os.MkdirAll(diretorio, 0755); err != nil {
		return fmt.Errorf(
			"erro ao criar diretório %s: %w", diretorio, err)
	}
	return os.WriteFile(caminho, dados, 0644)
}
