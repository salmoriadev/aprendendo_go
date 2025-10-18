package desafio01

import (
	"encoding/hex"
	"fmt"
)

func HexToString(hexString string) string {
	decodedBytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Erro ao decodificar a string hexadecimal:", err)
		return ""
	}
	resultString := string(decodedBytes)
	return resultString
}

func Desafio01() (string, error) {
	hexString := ("686974207468652062756c6c2773206" +
		"5796549276d206b696c6c696e6720796f757220627261696e206c696b" +
		"65206120706f69736f6e6f7573206d757368726f6f6d")
	resultString := HexToString(hexString)
	return resultString, nil
}
