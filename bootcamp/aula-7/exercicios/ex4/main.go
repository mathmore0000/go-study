package main

import (
	"fmt"
)

/*
Exercício 4 - Impostos sobre o salário #4

Repeti o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de erro receba como parâmetro o valor de "salary"

	indicando que ele não atinge o mínimo tributável
	(a mensagem exibida pelo console deve dizer:“Error: the minimum taxable amount is 150,000 and the salary entered is: [salary]”,
	sendo  [salary] o valor do tipo int passado pelo parâmetro).
*/
func main() {
	var salary int = 2000

	ok, err := isSalaryTaxable(salary)

	if err != nil {
		fmt.Println(err)
		return
	}

	if ok {
		fmt.Println("Salary is taxable")
	}
}

func isSalaryTaxable(salary int) (bool, error) {
	if salary <= 150000 {
		return false, fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: [%v]", salary)
	}
	return true, nil
}
