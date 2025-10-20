package desafio07

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

func Desafio07() (string, error) {
	chave := []byte("YELLOW SUBMARINE")

	textoCifradoB64, err := os.ReadFile("desafio07/7.txt")
	if err != nil {
		return "", fmt.Errorf("falha ao ler o arquivo 7.txt: %w", err)
	}

	textoCifrado, err := base64.StdEncoding.DecodeString(string(textoCifradoB64))
	if err != nil {
		return "", fmt.Errorf("falha ao decodificar base64: %w", err)
	}

	textoPlano, err := DecifrarAESECB(textoCifrado, chave)
	if err != nil {
		return "", err
	}

	return string(textoPlano), nil
}

func DecifrarAESECB(dadosCifrados, chave []byte) ([]byte, error) {
	cifra, err := aes.NewCipher(chave)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar a cifra AES: %w", err)
	}

	tamanhoDoBloco := cifra.BlockSize()
	if len(dadosCifrados)%tamanhoDoBloco != 0 {
		return nil, fmt.Errorf("o texto cifrado não é um múltiplo do tamanho do bloco")
	}

	textoPlano := make([]byte, len(dadosCifrados))

	for i := 0; i < len(dadosCifrados); i += tamanhoDoBloco {
		cifra.Decrypt(textoPlano[i:i+tamanhoDoBloco], dadosCifrados[i:i+tamanhoDoBloco])
	}
	textoPlano, err = validarERemoverPadding(textoPlano)
	if err != nil {
		return nil, fmt.Errorf("falha ao validar/remover o padding: %w", err)
	}

	return textoPlano, nil
}

func validarERemoverPadding(dados []byte) ([]byte, error) {
	tamanho := len(dados)
	if tamanho == 0 {
		return nil, fmt.Errorf("dados de entrada vazios")
	}

	tamanhoDoPadding := int(dados[tamanho-1])
	if tamanhoDoPadding == 0 || tamanhoDoPadding > tamanho {
		return nil, fmt.Errorf("valor de padding inválido: %d", tamanhoDoPadding)
	}

	paddingEsperado := bytes.Repeat([]byte{byte(tamanhoDoPadding)}, tamanhoDoPadding)
	if !bytes.HasSuffix(dados, paddingEsperado) {
		return nil, fmt.Errorf("padding inconsistente")
	}

	return dados[:(tamanho - tamanhoDoPadding)], nil
}
