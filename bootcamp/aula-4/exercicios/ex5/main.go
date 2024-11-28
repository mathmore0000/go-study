package qtdRacao

import (
	"errors"
	"fmt"
)

/*
	Exercício 5 - Calcular a quantidade de alimentos

	Um abrigo de animais precisa calcular a quantidade de alimentos que precisa comprar para seus animais de estimação.
	No momento, eles só têm tarântulas, hamsters, cães e gatos, mas esperam poder abrigar muito mais animais. Eles têm os seguintes dados:

	Cão 10 kg de comida.
	Gato 5 kg de comida.
	Hamster 250 g de comida.
	Tarântula 150 g de comida.
	É solicitado que você:

	Implemente uma função Animal que receba como parâmetro um valor de texto com o animal especificado e retorne uma função e uma mensagem (caso o animal não exista).
	Uma função para cada animal que calcule a quantidade de comida com base na quantidade do tipo de animal especificado.
	Exemplo:

	const (
	dog    = "dog"
	cat    = "cat"
	)
	...
	animalDog, msg := animal(dog)
	animalCat, msg := animal(cat)
	...
	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
*/

const (
	dog = "dog"
	cat = "cat"
)

func main() {
	animalDog, msg := animal(dog)
	if msg != "" {
		panic(errors.New(msg))
	}
	animalCat, msg := animal(cat)
	if msg != "" {
		panic(errors.New(msg))
	}

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)

	fmt.Println("Quantidade total de ração ->", amount)
}
func animal(animal string) (func(qtd int) float64, string) {
	switch animal {
	case "dog":
		return getCachorroComida, ""
	case "cat":
		return getGatoComida, ""
	case "hasmter":
		return getHamsterComida, ""
	case "tarantula":
		return getTarantulaComida, ""
	}

	return getCachorroComida, "Animal não encontrado"
}

func getCachorroComida(qtd int) float64 {
	if qtd < 0 {
		return 0
	}
	return float64(qtd) * 10
}
func getTarantulaComida(qtd int) float64 {
	if qtd < 0 {
		return 0
	}
	return float64(qtd) * .150
}

func getGatoComida(qtd int) float64 {
	if qtd < 0 {
		return 0
	}
	return float64(qtd) * 5
}

func getHamsterComida(qtd int) float64 {
	if qtd < 0 {
		return 0
	}
	return float64(qtd) * .250
}
