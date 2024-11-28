package main

import "fmt"

/*
	Exercício 3 - Calcular o salário

	Uma empresa marítima precisa calcular o salário de seus funcionários com base no número de horas trabalhadas por mês e na categoria.

	Categoria C, o salário deles é de US$ 1.000 por hora.
	Categoria B, o salário deles é de US$ 1.500 por hora, mais 20% do salário mensal.
	Categoria A, o salário deles é de US$ 3.000 por hora, mais 50% do salário mensal.
	Você deve gerar uma função que receba como parâmetro o número de minutos trabalhados por mês, a categoria e retorne o salário.
*/

func main() {
	salario := getSalario(44, "C")

	fmt.Printf("A salário desse colaborador é USD $%v\n", salario)
}

func getSalario(horas int, categoria string) (salario float32) {
	switch categoria {
	case "A":
		return getSalarioCatA(horas)
	case "B":
		return getSalarioCatB(horas)
	case "C":
		return getSalarioCatC(horas)
	}
	return
}

func getSalarioCatA(horas int) (salario float32) {
	salario = float32(horas * 1000)
	return
}

func getSalarioCatB(horas int) (salario float32) {
	salario = float32(horas * 1500)
	salario = salario + (salario * 0.2)
	return
}

func getSalarioCatC(horas int) (salario float32) {
	salario = float32(horas * 3000)
	salario = salario + (salario * 0.5)
	return
}
