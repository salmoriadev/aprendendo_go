package desafio07

/*
Decifrador AES ECB (Desafio 07)

Este arquivo lê o texto base64 fornecido pelo desafio 7, decodifica o conteúdo
com a chave fixa "YELLOW SUBMARINE" e devolve o texto em claro. O objetivo é
apenas demonstrar a decifragem ECB do conjunto Cryptopals Set 1, sem
generalizações extras.
*/

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

/*
Desafio07 carrega o arquivo `desafio07/7.txt`, decodifica de base64 e aplica a
decifragem AES-128 em ECB. O ciphertext do desafio já está alinhado aos blocos
de 16 bytes, então o resultado é retornado exatamente como sai do AES sem
validações adicionais de padding.
*/
func Desafio07() (string, error) {
	chave := []byte("YELLOW SUBMARINE")

	textoCifradoB64, err := os.ReadFile("desafio07/7.txt")
	if err != nil {
		return "", fmt.Errorf("falha ao ler o arquivo 7.txt: %w", err)
	}

	textoCifrado, err := base64.StdEncoding.DecodeString(
		string(textoCifradoB64))
	if err != nil {
		return "", fmt.Errorf("falha ao decodificar base64: %w", err)
	}

	textoPlano, err := DecifrarAESECB(textoCifrado, chave)
	if err != nil {
		return "", err
	}

	return string(textoPlano), nil
}

/*
DecifrarAESECB aplica AES-128 ECB direto sobre os blocos do Ciphertext. O Set 1
garante que o tamanho seja múltiplo de 16, portanto retornamos o buffer de
plaintext sem mexer em padding.
*/
func DecifrarAESECB(dadosCifrados, chave []byte) ([]byte, error) {
	cifra, err := aes.NewCipher(chave)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar a cifra AES: %w", err)
	}

	tamanhoDoBloco := cifra.BlockSize()
	if len(dadosCifrados)%tamanhoDoBloco != 0 {
		return nil, fmt.Errorf(
			"o texto cifrado não é um múltiplo do tamanho do bloco")
	}

	textoPlano := make([]byte, len(dadosCifrados))

	for i := 0; i < len(dadosCifrados); i += tamanhoDoBloco {
		cifra.Decrypt(
			textoPlano[i:i+tamanhoDoBloco],
			dadosCifrados[i:i+tamanhoDoBloco])
	}

	return textoPlano, nil
}
