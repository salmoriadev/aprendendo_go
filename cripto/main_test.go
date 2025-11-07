package main

import "testing"

// TestNovaConfiguracaoPadrao garante que a configuração padrão mantém os valores esperados.
func TestNovaConfiguracaoPadrao(t *testing.T) {
	config := novaConfiguracaoPadrao()

	if config.TamanhoChaveBits != 2048 {
		t.Fatalf("tamanho padrão da chave deveria ser 2048, obtido %d",
			config.TamanhoChaveBits)
	}
	if config.ValidadeCertificadoACAnos != 10 {
		t.Fatalf("validade padrão da AC deveria ser 10 anos, obtido %d",
			config.ValidadeCertificadoACAnos)
	}
	if config.ValidadeCertificadoAnos != 1 {
		t.Fatalf("validade padrão do certificado de usuário deveria ser 1 ano, obtido %d",
			config.ValidadeCertificadoAnos)
	}

	if config.EstrategiaChave == nil || config.EstrategiaCertificado == nil {
		t.Fatalf("estratégias não podem ser nulas na configuração padrão")
	}
	if config.EstrategiaResumo == nil || config.EstrategiaAssinatura == nil {
		t.Fatalf("estratégias de resumo e assinatura devem estar definidas")
	}
}
