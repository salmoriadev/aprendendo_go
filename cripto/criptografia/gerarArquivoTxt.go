package criptografia

import (
	"fmt"
	"os"
	"path/filepath"
)

func GerarArquivoTexto(caminho string, conteudo string) error {
	diretorio := filepath.Dir(caminho)

	if err := os.MkdirAll(diretorio, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório %s: %w", diretorio, err)
	}

	err := os.WriteFile(caminho, []byte(conteudo), 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever arquivo: %w", err)
	}

	return nil
}
