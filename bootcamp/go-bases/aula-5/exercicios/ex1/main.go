package main

import (
	"errors"
	"fmt"
)

/*
	Exercício 1
	Crie um programa que atenda aos seguintes pontos:

	Ter uma estrutura chamada Product com os campos ID, Name, Price, Description e Category.
	Ter uma fatia global de Produto chamada Produtos instanciada com valores. 2 métodos associados à estrutura Produto: Save(), GetAll().
		O método Save() deve pegar a fatia de Products e adicionar o produto a partir do qual o método é chamado.
		O método GetAll() imprime todos os produtos salvos na fatia Products.
	Uma função getById() para a qual um INT deve ser passado como parâmetro e retorna o produto correspondente ao parâmetro passado.
	Execute pelo menos uma vez cada método e função definidos em main().
*/

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}
type ProductList []Product

var products ProductList = ProductList{
	Product{ID: 1, Name: "Produto 1", Price: 99.99, Description: "Produto 1 é definitivamente um produto", Category: "Produtos"},
	Product{ID: 2, Name: "Produto 2", Price: 9.99, Description: "Produto 2 é definitivamente um produto", Category: "Produtos"},
	Product{ID: 3, Name: "Produto 3", Price: 149.99, Description: "Produto 3 é definitivamente um produto", Category: "Produtos"},
}

func (product Product) Save() {
	products = append(products, product)
}

func getById(id int) (Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("Produto não encontrado")
}

func (p ProductList) GetAll() {
	for _, product := range p {
		fmt.Printf("ID -> %v | Name -> %v | Price -> %v | Description -> %v | Category -> %v\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func main() {
	fmt.Println("Produtos iniciais:")
	products.GetAll()

	newProduct := Product{ID: 4, Name: "Produto 4", Price: 29.99, Description: "Produto 4 é um novo produto", Category: "Novos Produtos"}
	newProduct.Save() // Salvando o novo produto

	fmt.Println("Produtos alterados:")
	products.GetAll()

	fmt.Println("\nBuscando Produto com ID 2:")
	product, err := getById(4)
	if err == nil {
		fmt.Printf("Produto encontrado: ID -> %v | Name -> %v | Price -> %v | Description -> %v | Category -> %v\n",
			product.ID, product.Name, product.Price, product.Description, product.Category)
	} else {
		fmt.Println(err)
	}
}
