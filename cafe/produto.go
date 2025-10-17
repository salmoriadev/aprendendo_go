package main

import "fmt"

type Produto interface {
	CalcularLucro() float64
	VerAtributos() string
}

type ProdutoBase struct {
	Nome        string
	PrecoVenda  float64
	PrecoCompra float64
}

func (p *ProdutoBase) CalcularLucro() float64 {
	return p.PrecoVenda - p.PrecoCompra
}

func (p *ProdutoBase) VerAtributos() string {
	return fmt.Sprintf("Nome: %s, Preço de Venda: %.2f, Preço de Compra: %.2f",
		p.Nome, p.PrecoVenda, p.PrecoCompra)
}
