package criptografia

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func TxtParaPdf(caminhoTxt string, caminhoPdf string) {
	text, err := os.ReadFile(caminhoTxt)
	if err != nil {
		fmt.Println("Erro ao ler arquivo TXT:", err)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.MoveTo(0, 10)
	pdf.Cell(1, 1, "Lorem Ipsum")
	pdf.MoveTo(0, 20)
	pdf.SetFont("Arial", "", 14)
	width, _ := pdf.GetPageSize()
	pdf.MultiCell(width, 10, string(text), "", "", false)
	err = pdf.OutputFileAndClose(caminhoPdf)
	if err == nil {
		fmt.Println("PDF gerado com sucesso")
	}
}
