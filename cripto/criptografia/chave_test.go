package criptografia

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

// TestConversoesPEM valida que as conversões de chave para PEM retornam dados não vazios.
func TestConversoesPEM(t *testing.T) {
	chavePrivada, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("falha ao gerar chave RSA de teste: %v", err)
	}

	pemPrivada := ChavePrivadaParaPEM(chavePrivada)
	if len(pemPrivada) == 0 {
		t.Fatalf("conversão da chave privada para PEM não pode retornar vazio")
	}

	pemPublica := ChavePublicaParaPEM(&chavePrivada.PublicKey)
	if len(pemPublica) == 0 {
		t.Fatalf("conversão da chave pública para PEM não pode retornar vazio")
	}
}

// TestResumoSha256 verifica se a estratégia de resumo produz saída consistente.
func TestResumoSha256(t *testing.T) {
	estrategia := NovaEstrategiaResumoSha256()
	dados := []byte("LabSEC Portfolio")
	resumo1 := estrategia.Resumir(dados)
	resumo2 := estrategia.Resumir(dados)

	if len(resumo1) == 0 {
		t.Fatalf("hash não deveria ser vazio")
	}
	if string(resumo1) != string(resumo2) {
		t.Fatalf("hash deve ser determinístico para a mesma entrada")
	}

	if estrategia.HashFunc() != crypto.SHA256 {
		t.Fatalf("HashFunc deve retornar crypto.SHA256")
	}
}
