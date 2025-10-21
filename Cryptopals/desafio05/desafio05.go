package desafio05

/*
Desafio 05 - Cifra Repetida XOR
Neste desafio, você precisa cifrar uma string usando uma cifra XOR
repetida. A cifra XOR repetida funciona aplicando a operação XOR entre
cada byte da mensagem e um byte correspondente da chave, repetindo a chave
conforme necessário para cobrir toda a mensagem.
O código reutiliza a função de XOR do desafio 2.
*/

import (
	"cryptopals/desafio02"
	"encoding/hex"
)

func Desafio05() (string, error) {
	mensagem := "Burning 'em, if you ain't quick and nimble \n" +
		"I go crazy when I hear a cymbal"
	chave := "ICE"
	bytesMensagem := ([]byte)(mensagem)
	bytesChave := ([]byte)(chave)
	mensagem_cifrada := hex.EncodeToString(
		desafio02.XorBytes(bytesMensagem, bytesChave))
	return mensagem_cifrada, nil

}
