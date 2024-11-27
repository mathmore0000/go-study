package main

import "fmt"

/*
	Exercício 3 - Qual mês corresponde a

	Crie um aplicativo que receba uma variável com o número do mês.

	Dependendo do número, imprima o mês correspondente no texto.
	Você consegue pensar em maneiras diferentes de resolver isso? Qual delas você escolheria e por quê?
	Ex: 7, Julio.

	👀Observação: verifique se o número do mês está correto.
*/

func main() {
	meses := map[int]string{1: "Janeiro", 2: "Fevereiro", 3: "Março", 4: "Abril", 5: "Maio", 6: "Junho", 7: "Julho", 8: "Agosto", 9: "Setembro", 10: "Outubro", 11: "Novembro", 12: "Dezembro"}
	var mes int = 9

	fmt.Println("Mês escolhido ->", meses[mes])
}
