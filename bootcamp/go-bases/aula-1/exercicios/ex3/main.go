package main

import "fmt"

/*
	Exercício 3 - Declaração de variáveis
	Um professor de programação está corrigindo os exames de seus alunos na disciplina Programação I para dar a eles os retornos correspondentes. Um dos itens do exame é declarar diferentes variáveis.
	Preciso de ajuda para:
	1. Detecte quais dessas variáveis declaradas pelo aluno estão corretas.
	2. Corrija as incorretas.
	var 1firstName string
	var lastName string
	var int age
	1lastName := 6
	var driver_license = true
	var person height int
	childsNumber := 2
*/

func main() {
	var firstName string
	var lastName string
	var age int
	lastName = "Ultimo nome"
	var driverLicense bool = true
	var height float32
	childsNumber := 2

	fmt.Println(firstName, lastName, age, driverLicense, height, childsNumber)
}
