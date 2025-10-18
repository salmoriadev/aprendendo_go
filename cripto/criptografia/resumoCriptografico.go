package criptografia

import (
	"crypto/sha256"
	"os"
)

func GerarResumoCriptografico(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func PdfParaBytes(caminhoArquivoPdf string) ([]byte, error) {
	arrayBytes, err := os.ReadFile(caminhoArquivoPdf)
	if err != nil {
		return nil, err
	}
	return arrayBytes, nil
}
