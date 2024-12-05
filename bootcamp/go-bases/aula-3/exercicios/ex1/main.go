package main

import "fmt"

/*
	Exercício 1 - Impostos sobre o salário
	Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de depositar o salário.
	Para cumprir o objetivo, é necessário criar uma função que retorne o imposto de um salário.

	Levando em conta que se a pessoa ganha mais de US$ 50.000, 17% do salário será deduzido e se a pessoa ganha mais de US$ 150.000,
	10% também será deduzido (27% no total).
*/

func main() {
	var salario float32 = 180000
	imposto := getImposto(salario)

	fmt.Printf("Imposto para um salário de BRL %v é de %v%\n", salario, imposto)
}

func getImposto(salario float32) (imposto float32) {
	if salario > 150000 {
		return 27
	}
	if salario > 50000 {
		return 17
	}
	return
}
