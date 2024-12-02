package main

import (
	"errors"
	"fmt"
)

/*
	Exercício 3 - Impostos sobre o salário #3

	Faça o mesmo que no exercício anterior, mas reformule o código de modo que, em vez de "Error()", seja implementado "errors.New()".
*/

type errSalary struct {
	message string
}

var ErrSalaryTooLow = errors.New("salary is less than 10000")

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
