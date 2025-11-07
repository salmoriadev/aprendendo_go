package desafio05

/*
Cifra XOR Repetida (Desafio 05)

Aplica a cifra repetida para a frase do desafio 5 usando a função `XorBytes`
do desafio 2, retornando o resultado em hex como pedido pelo Set 1.
*/

import (
	"cryptopals/desafio02"
	"encoding/hex"
)

/*
Desafio05 apenas instancia mensagem e chave exigidas pelo enunciado e retorna a
Cifra XOR repetida em hex.
*/
func Desafio05() (string, error) {
	mensagem := "Burning 'em, if you ain't quick and nimble \n" +
		"I go crazy when I hear a cymbal"
	chave := "ICE"
	bytesMensagem := ([]byte)(mensagem)
	bytesChave := ([]byte)(chave)
	mensagemCifrada := hex.EncodeToString(
		desafio02.XorBytes(bytesMensagem, bytesChave))
	return mensagemCifrada, nil
}
