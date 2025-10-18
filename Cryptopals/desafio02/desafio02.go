package desafio02

import (
	"encoding/hex"
)

func XorBytes(mensagem, chave []byte) []byte {
	if len(mensagem) > len(chave) {
		tamanhoChave := len(chave)
		for i := 0; i < len(mensagem); i++ {
			mensagem[i] ^= chave[i%tamanhoChave]
		}
	} else {
		for i := 0; i < len(mensagem); i++ {
			mensagem[i] ^= chave[i]
		}
	}
	return mensagem
}

func Desafio02() (string, error) {
	hexString1 := "1c0111001f010100061a024b53535009181c"
	hexString2 := "686974207468652062756c6c277320657965"

	decodedBytes1, err := hex.DecodeString(hexString1)
	if err != nil {
		return "", err
	}
	decodedBytes2, err := hex.DecodeString(hexString2)
	if err != nil {
		return "", err
	}
	result := XorBytes(decodedBytes1, decodedBytes2)
	resultString := hex.EncodeToString(result)
	return resultString, nil
}
