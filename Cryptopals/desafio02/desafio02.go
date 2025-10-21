package desafio02

/*
Desafio 02 - Operação XOR entre duas strings hexadecimais
Neste desafio, você precisa realizar uma operação XOR entre
duas strings hexadecimais. Apenas precisei fazer uma função
para fazer o XOR entre dois arrays de bytes dentro de um for loop.
Esse código será reutilizado no desafio 5.
*/

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
