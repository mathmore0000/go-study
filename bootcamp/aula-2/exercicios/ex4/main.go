package main

import "fmt"

/*
	Exercício 4 - Quantos anos tem...

	Um funcionário de uma empresa quer saber o nome e a idade de um de seus funcionários. De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	Por outro lado, também é necessário:

	Saber quantos de seus funcionários têm mais de 21 anos.
	Adicionar um novo funcionário à lista, chamado Federico, que tem 25 anos.
	Remover Pedro do mapa.
*/

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var qtdMais21 int

	for _, idade := range employees {
		if idade > 21 {
			qtdMais21++
		}

	}

	fmt.Println("Quantidade de usuários que têm mais de 21 anos ->", qtdMais21)
	employees["Frederico"] = 27
	delete(employees, "Pedro")
	fmt.Println(employees)
}
