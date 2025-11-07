/*
Testes que percorrem o fluxo completo de manipulação de arquivos (TXT, PDF,
hash e assinatura) garantindo a robustez dos serviços de suporte.
*/
package servicos

import (
	"path/filepath"
	"testing"

	"cripto/criptografia"
)

// TestFluxoCompletoComPDF cobre o fluxo de geração de PDF, resumo, assinatura e verificação.
func TestFluxoCompletoComPDF(t *testing.T) {
	t.Helper()
	diretorio := t.TempDir()

	caminhoTxt := filepath.Join(diretorio, "mensagem.txt")
	caminhoPdf := filepath.Join(diretorio, "mensagem.pdf")
	caminhoResumo := filepath.Join(diretorio, "resumo.txt")
	caminhoAssinatura := filepath.Join(diretorio, "assinatura.txt")

	err := GerarArquivoPDF(caminhoTxt, caminhoPdf, "conteudo de teste do LabSEC")
	if err != nil {
		t.Fatalf("GerarArquivoPDF falhou: %v", err)
	}

	estrategiaResumo := criptografia.NovaEstrategiaResumoSha256()
	resumo, err := ResumirPDF(caminhoPdf, caminhoResumo, estrategiaResumo)
	if err != nil {
		t.Fatalf("ResumirPDF falhou: %v", err)
	}
	if resumo == "" {
		t.Fatalf("resumo nao pode ser vazio")
	}

	estrategiaChave := criptografia.NovaEstrategiaChaveRSA()
	pares, err := estrategiaChave.GerarChavePrivada(1024)
	if err != nil {
		t.Fatalf("falha ao gerar chaves RSA para teste: %v", err)
	}

	estrategiaAssinatura := criptografia.NovaEstrategiaAssinaturaPkcs1v15()
	assinatura, err := AssinarDocumentoPDF(caminhoPdf, caminhoAssinatura,
		&pares, estrategiaResumo, estrategiaAssinatura)
	if err != nil {
		t.Fatalf("AssinarDocumentoPDF falhou: %v", err)
	}
	if len(assinatura) == 0 {
		t.Fatalf("assinatura nao pode ser vazia")
	}

	valida, err := VerificarAssinaturaDocumentoPDF(caminhoPdf, assinatura,
		pares.ChavePublica, estrategiaResumo, estrategiaAssinatura)
	if err != nil {
		t.Fatalf("VerificarAssinaturaDocumentoPDF retornou erro: %v", err)
	}
	if !valida {
		t.Fatalf("assinatura deveria ser valida para o proprio documento")
	}
}
