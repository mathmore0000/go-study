package main

import (
	"errors"
	"fmt"
)

/*
	Exercício 5 -  Impostos sobre o salário #5

	Vamos tornar nosso programa um pouco mais complexo e útil.

	Desenvolver as funções necessárias para permitir que a empresa faça cálculos:
	Salário mensal de um trabalhador de acordo com o número de horas trabalhadas. A função deve receber as horas trabalhadas no mês e o valor da hora como argumento.
	A função deve retornar mais de um valor (salário calculado e erro). Caso o salário mensal seja igual ou superior a US$ 150.000, será deduzido um imposto de 10%.
		Caso o número de horas mensais inserido seja inferior a 80 ou um número negativo, a função deve retornar um erro.
		Ela indicará “Error: the worker cannot have worked less than 80 hours per month”.

*/

func getMonthlySalary(hoursWorked int, hourValue float32) (float32, error) {
	if hoursWorked < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}
	var monthlySalary float32 = hourValue * float32(hoursWorked)

	if monthlySalary >= 150000 {
		return (monthlySalary - (monthlySalary * 0.1)), nil
	}
	return monthlySalary, nil
}

func main() {
	monthlySalary, err := getMonthlySalary(200, 1500)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Final salary ->", monthlySalary)
}
