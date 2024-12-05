package main

import (
	"fmt"
)

/*
	Exercício 1 - Impostos sobre o salário #1

	Em sua função "main", defina uma variável chamada "salary" e atribua a ela um valor do tipo "int".

	Crie um erro personalizado com uma estrutura que implemente "Error()" com a mensagem "Error: the salary entered does not reach the taxable minimum"
	e acione-a caso "salary" seja menor que 150.000. Caso contrário, será necessário imprimir no console a mensagem "Must pay tax".
*/

type ErrSalary struct {
	message string
}

func (e *ErrSalary) Error() string {
	return e.message
}

func main() {
	var salary int = 250000
	if salary < 150000 {

		err := &ErrSalary{message: "Error: the salary entered does not reach the taxable minimum"}
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Must pay tax")
}
