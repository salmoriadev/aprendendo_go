package desafio05

import (
	"cryptopals/desafio02"
	"encoding/hex"
	"fmt"
)

func Desafio05() {
	mensagem := "Burning 'em, if you ain't quick and nimble\n" +
		"I go crazy when I hear a cymbal"
	chave := "ICE"
	bytesMensagem := ([]byte)(mensagem)
	bytesChave := ([]byte)(chave)
	mensagem_cifrada := hex.EncodeToString(desafio02.XorBytes(bytesMensagem, bytesChave))
	fmt.Println(mensagem_cifrada)

}
