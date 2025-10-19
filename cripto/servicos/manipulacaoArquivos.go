package servicos

import (
	"cripto/criptografia"
	"crypto/rsa"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

func GerarArquivoPDF(caminhoArquivoTxt string, caminhoArquivoPdf string,
	conteudoArquivo string) {
	err := gerarArquivoTexto(caminhoArquivoTxt, conteudoArquivo)
	if err != nil {
		log.Fatalf("Erro ao gerar .txt: %v", err)
	}
	txtParaPdf(caminhoArquivoTxt, caminhoArquivoPdf)
}

func ResumirPDF(caminhoArquivoPdf string, caminhoResumo string,
	estrategiaResumo criptografia.EstrategiaResumo) {
	dados, err := pdfParaBytes(caminhoArquivoPdf)
	if err != nil {
		log.Fatalf("Erro ao ler PDF: %v", err)
	}
	resumo := hex.EncodeToString(estrategiaResumo.Resumir(dados))

	fmt.Println("Resumo criptográfico do PDF:", resumo)
	err = gerarArquivoTexto(caminhoResumo, resumo)
	if err != nil {
		log.Fatalf("Erro ao gerar resumo: %v", err)
	}
}

func AssinarDocumentoPDF(caminhoArquivoPdf string, caminhoAssinatura string,
	chaves criptografia.ParDeChaves, estrategiaResumo criptografia.EstrategiaResumo,
	estrategiaAssinatura criptografia.EstrategiaAssinatura) []byte {

	dados, err := pdfParaBytes(caminhoArquivoPdf)
	if err != nil {
		log.Fatalf("Erro ao ler PDF: %v", err)
	}

	resumoEmBytes := estrategiaResumo.Resumir(dados)
	assinatura, err := estrategiaAssinatura.Assinar(resumoEmBytes, chaves.ChavePrivada)
	if err != nil {
		log.Fatalf("Erro ao assinar o resumo: %v", err)
	}

	err = gerarArquivoTexto(caminhoAssinatura, hex.EncodeToString(assinatura))
	if err != nil {
		log.Fatalf("Erro ao gerar assinatura: %v", err)
	}
	return assinatura
}

func VerificarAssinaturaDocumentoPDF(caminhoArquivoPdf string,
	assinatura []byte, chavePublica *rsa.PublicKey,
	estrategiaResumo criptografia.EstrategiaResumo,
	estrategiaAssinatura criptografia.EstrategiaAssinatura) bool {

	dados, err := pdfParaBytes(caminhoArquivoPdf)
	if err != nil {
		log.Fatalf("Erro ao ler PDF: %v", err)
	}

	resumoEmBytes := estrategiaResumo.Resumir(dados)
	err = estrategiaAssinatura.VerificarAssinatura(resumoEmBytes,
		assinatura, chavePublica)
	if err != nil {
		log.Printf("Erro ao verificar a assinatura: %v", err)
		return false
	}
	return true
}

func gerarArquivoTexto(caminho string, conteudo string) error {
	diretorio := filepath.Dir(caminho)
	if err := os.MkdirAll(diretorio, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório %s: %w", diretorio, err)
	}
	return os.WriteFile(caminho, []byte(conteudo), 0644)
}

func pdfParaBytes(caminhoArquivoPdf string) ([]byte, error) {
	return os.ReadFile(caminhoArquivoPdf)
}

func txtParaPdf(caminhoTxt string, caminhoPdf string) {
	text, err := os.ReadFile(caminhoTxt)
	if err != nil {
		fmt.Println("Erro ao ler arquivo TXT:", err)
		return
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 14)
	width, _ := pdf.GetPageSize()
	pdf.MultiCell(width, 10, string(text), "", "", false)
	err = pdf.OutputFileAndClose(caminhoPdf)
	if err == nil {
		fmt.Println("PDF gerado com sucesso")
	} else {
		log.Printf("Erro ao gerar PDF: %v", err)
	}
}

func escreverArquivo(caminho string, dados []byte) error {
	diretorio := filepath.Dir(caminho)
	if err := os.MkdirAll(diretorio, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório %s: %w", diretorio, err)
	}
	return os.WriteFile(caminho, dados, 0644)
}
