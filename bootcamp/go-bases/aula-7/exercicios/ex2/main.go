package main

import (
	"errors"
	"fmt"
)

/*
	Exercício 2 - Impostos sobre o salário #2

	Em sua função "main", defina uma variável chamada "salary" e atribua a ela um valor do tipo "int".
		Crie um erro personalizado com uma estrutura que implemente "Error()" com a mensagem "Error: salary is less than 10000" e a inicie caso "salary" seja menor ou igual a 10000.
		A validação deve ser feita com a função "Is()" dentro do "main".
*/

type errSalary struct {
	message string
}

func (e *errSalary) Error() string {
	return "Error: " + e.message
}

var ErrSalaryTooLow = &errSalary{message: "salary is less than 10000"}

func main() {
	var salary int = 2000

	ok, err := isSalaryOk(salary)

	if err != nil {
		if errors.Is(err, ErrSalaryTooLow) {
			fmt.Println(err)
			return
		}
		fmt.Println("Erro genérico")
	}

	if ok {
		fmt.Println("Salary ok")
	}
}

func isSalaryOk(salary int) (bool, error) {
	if salary <= 10000 {
		return false, ErrSalaryTooLow
	}
	return true, nil
}
