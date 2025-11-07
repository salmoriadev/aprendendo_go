/*
Testes que validam a geração de pares de chaves e certificados utilizando
diretórios temporários para comprovar o funcionamento dos serviços de
orquestração.
*/
package servicos

import (
	"os"
	"path/filepath"
	"testing"

	"cripto/criptografia"
)

// TestGerarParDeChavesEGerarCertificados garante que as chaves e os
// certificados são criados sem erros.
func TestGerarParDeChavesEGerarCertificados(t *testing.T) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic inesperado: %v", r)
		}
	}()

	diretorio := t.TempDir()
	estrategiaChave := criptografia.NovaEstrategiaChaveRSA()

	parAutoridade, err := GerarParDeChaves(1024, diretorio, "autoridade_teste", estrategiaChave)
	if err != nil {
		t.Fatalf("erro ao gerar chaves da autoridade: %v", err)
	}

	parUsuario, err := GerarParDeChaves(1024, diretorio, "usuario_teste", estrategiaChave)
	if err != nil {
		t.Fatalf("erro ao gerar chaves do usuario: %v", err)
	}

	estrategiaCert := criptografia.NovaEstrategiaCertificado()
	dadosAutoridade := DadosIdentificacaoCertificado{
		Organizacao: "LabSEC",
		Pais:        "BR",
		Provincia:   "Santa Catarina",
		Localidade:  "Florianopolis",
		NomeComum:   "LabSEC Autoridade",
	}

	dadosUsuario := DadosIdentificacaoCertificado{
		Organizacao: "LabSEC",
		Pais:        "BR",
		Provincia:   "Santa Catarina",
		Localidade:  "Florianopolis",
		NomeComum:   "usuario.teste",
	}

	if err := GerarCertificados(parAutoridade, parUsuario,
		5, 1, diretorio, dadosAutoridade, dadosUsuario, estrategiaCert); err != nil {
		t.Fatalf("erro ao gerar certificados: %v", err)
	}

	verificarArquivoExiste(t, filepath.Join(diretorio, "chave_autoridade_teste_privada.pem"))
	verificarArquivoExiste(t, filepath.Join(diretorio, "chave_autoridade_teste_publica.pem"))
	verificarArquivoExiste(t, filepath.Join(diretorio, "chave_usuario_teste_privada.pem"))
	verificarArquivoExiste(t, filepath.Join(diretorio, "chave_usuario_teste_publica.pem"))
	verificarArquivoExiste(t, filepath.Join(diretorio, "certificado_autoridade.pem"))
	verificarArquivoExiste(t, filepath.Join(diretorio, "certificado_usuario.pem"))

	if parAutoridade.ChavePrivada == nil || parUsuario.ChavePrivada == nil {
		t.Fatalf("pares de chaves nao deveriam ser nulos")
	}
}

// verificarArquivoExiste falha o teste se o arquivo esperado não estiver no diretório.
func verificarArquivoExiste(t *testing.T, caminho string) {
	t.Helper()
	if _, err := os.Stat(caminho); err != nil {
		t.Fatalf("arquivo esperado nao encontrado: %s (%v)", caminho, err)
	}
}
