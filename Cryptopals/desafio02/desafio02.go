package desafio02

/*
Operações XOR (Desafio 02)

Resolve o desafio 2 do Set 1: aplicar XOR posição a posição em dois buffers
hexadecimais e devolver o resultado em hex. O helper é reaproveitado pelo
desafio 5.
*/

import (
	"encoding/hex"
)

/*
XorBytes executa XOR byte a byte retornando um novo slice, repetindo a chave
quando necessário. Mantemos a função sem efeitos colaterais para reutilizar nos
demais desafios do Set 1.
*/
func XorBytes(mensagem, chave []byte) []byte {
	saida := make([]byte, len(mensagem))
	tamanhoChave := len(chave)
	for i := 0; i < len(mensagem); i++ {
		saida[i] = mensagem[i] ^ chave[i%tamanhoChave]
	}
	return saida
}

/*
Desafio02 usa XorBytes em duas strings fixas do enunciado e retorna a resposta
em hex.
*/
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
