package main

import "fmt"

/*
	Um banco deseja conceder empréstimos a seus clientes, mas nem todos têm acesso a eles. Ele tem certas regras para determinar quais clientes podem receber empréstimos.
	O banco só concede empréstimos a clientes com mais de 22 anos de idade, que estejam empregados e que estejam em seu emprego há mais de um ano.
	Dentro dos empréstimos concedidos, o banco não cobrará juros daqueles que tiverem um salário superior a US$ 100.000.
	Você precisa criar um aplicativo que receba essas variáveis e imprima uma mensagem de acordo com cada caso.
	📌Dica: seu código deve ser capaz de imprimir pelo menos três mensagens diferentes.
*/

func main() {
	var valorFinalEmprestimo float32
	var valorEmprestimo float32
	var idadeCliente int8
	var salarioCliente float32
	var tempoEmprego int8 // Em mês
	var juros int8 = 10

	valorEmprestimo = 10000
	salarioCliente = 50000
	idadeCliente = 18
	tempoEmprego = 24

	if idadeCliente < 22 || tempoEmprego < 12 {
		fmt.Println("Seu empréstimo não foi aceito, :(")
		return
	}

	if salarioCliente < 100000 {
		valorFinalEmprestimo = valorEmprestimo + (valorEmprestimo * (float32(juros) / 100))
		fmt.Println("Seu empréstimo foi aceito, valor final ->", valorFinalEmprestimo)
		return
	}

	fmt.Println("Seu empréstimo foi aceito, e sem juros!")

}
