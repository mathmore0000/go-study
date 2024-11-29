package main

import (
	"errors"
	"fmt"
)

/*
	Exercício 2 - Produtos

	Algumas lojas de comércio eletrônico precisam criar uma funcionalidade no Go para gerenciar produtos e retornar o valor do preço total.
		A empresa tem três tipos de produtos: Pequeno, Médio e Grande (muitos outros são esperados).

	E os custos adicionais são:

	Pequeno: apenas o custo do produto
	Médio: o preço do produto + 3% do produto + 3% de mantê-lo na loja
	Grande: o preço do produto + 6% de mantê-lo na loja e, além disso, o custo de envio de US$ 2.500.

	O custo de manter o produto em estoque na loja é uma porcentagem do preço do produto.

	É necessária uma função factory que receba o tipo de produto e o preço e retorne uma interface Product que tenha o método Price.

	Deve ser possível executar o método Price e fazer com que o método retorne o preço total com base no custo do produto e em quaisquer custos adicionais.
*/

type IProduct interface {
	Price() float32
}

type Pequeno struct {
	custo float32
}

type Medio struct {
	custo float32
}

type Grande struct {
	custo float32
}

func (p *Pequeno) Price() float32 {
	return p.custo
}

func (m *Medio) Price() float32 {
	return m.custo + (m.custo * 0.3)
}

func (g *Grande) Price() float32 {
	return g.custo + 2500
}

func factory(tipo string, custo float32) (IProduct, error) {
	switch tipo {
	case ("pequeno"):
		return &Pequeno{custo}, nil
	case ("medio"):
		return &Medio{custo}, nil
	case ("grande"):
		return &Grande{custo}, nil
	}
	return &Pequeno{}, errors.New("Tipo não encontrado")

}

func main() {
	p, err := factory("pequeno", 100.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	m, err := factory("medio", 200)
	if err != nil {
		fmt.Println(err)
		return
	}

	g, err := factory("grande", 500)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Pequeno price: %.2f\n", p.Price())
	fmt.Printf("Medio price: %.2f\n", m.Price())
	fmt.Printf("Grande price: %.2f\n", g.Price())
}
