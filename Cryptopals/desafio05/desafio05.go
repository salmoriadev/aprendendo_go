package desafio05

import (
	"cryptopals/desafio02"
	"encoding/hex"
)

func Desafio05() (string, error) {
	mensagem := "Burning 'em, if you ain't quick and nimble\n" +
		"I go crazy when I hear a cymbal"
	chave := "ICE"
	bytesMensagem := ([]byte)(mensagem)
	bytesChave := ([]byte)(chave)
	mensagem_cifrada := hex.EncodeToString(desafio02.XorBytes(bytesMensagem, bytesChave))
	return mensagem_cifrada, nil

}
