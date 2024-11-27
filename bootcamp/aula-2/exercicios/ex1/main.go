package main

import "fmt"

/*
	Exercício 1 - Letras em uma palavra
	A Real Academia Espanhola quer saber quantas letras uma palavra tem e, em seguida, ter cada uma das letras separadamente para soletrá-la. Para isso, eles terão de:
	Criar um aplicativo que receba uma variável com a palavra e imprima o número de letras que ela contém.
	Em seguida, imprima cada uma das letras.
*/

func main() {
	var palavra = "Academia"

	fmt.Println("Quantidade de letras ->", len(palavra))

	fmt.Println("\nLetras:")

	for _, letra := range palavra {
		fmt.Println(string(letra))
	}
}
