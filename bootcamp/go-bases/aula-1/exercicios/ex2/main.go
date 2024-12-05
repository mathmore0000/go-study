package main

import "fmt"

/*
	Exercício 2 - Clima
	Uma empresa de meteorologia deseja um aplicativo em que possa ter a temperatura, a umidade e a pressão atmosférica de diferentes locais.
		1

	1. Declare 3 variáveis especificando o tipo de dados; como valor, elas devem ter a temperatura, a umidade e a pressão do local onde você está.
	2. Imprima os valores das variáveis no console.
	3. Que tipo de dados você atribuiri a às variáveis?
*/

func main() {
	var temperatura float32 = 23.5
	var umidade float32 = 50
	var pressao int = 1000
	fmt.Println("A temperatura atual é", temperatura, "oº")
	fmt.Println("A umidade atual é", umidade, "%")
	fmt.Println("A pressao atual é", pressao, "hPa")
}
