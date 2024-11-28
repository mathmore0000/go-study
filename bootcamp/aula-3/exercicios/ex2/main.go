package main

import "fmt"

/*
	Exercício 2 - Calcular a média

	Uma escola precisa calcular a média (por aluno) de suas notas.
	É solicitado que ela gere uma função na qual possam ser passados N números inteiros e que retorne a média.
	Não é possível inserir notas negativas.
*/

func main() {
	media := getMedia(5, 4, 2, 3, 4.9, -2)

	fmt.Printf("A média desse aluno é %v%\n", media)
}

func getMedia(notas ...float32) (media float32) {
	var nota float32
	var lenNotas int = len(notas)
	for _, nota = range notas {
		if nota < 0 {
			lenNotas--
			continue
		}
		media += nota
	}
	return media / float32(lenNotas)
}
