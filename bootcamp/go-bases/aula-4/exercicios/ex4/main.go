package minMaxAvg

import (
	"errors"
	"fmt"
)

/*
	Exercício 4 - Cálculo de estatísticas

	Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de notas para os alunos de um curso.
	Para isso, eles precisam gerar uma função que indique o tipo de cálculo que desejam realizar (mínimo, máximo ou médio) e
	que retorne outra função e uma mensagem (caso o cálculo não esteja definido) que possa receber um número N de inteiros e
	retorne o cálculo indicado na função anterior. Exemplo:

	const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
	)
	...
	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)
	...
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		panic(errors.New(err.Error()))
	}

	averageFunc, err := operation(average)
	if err != nil {
		panic(errors.New(err.Error()))
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		panic(errors.New(err.Error()))
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("O minímo é %v\n", minValue)
	fmt.Printf("A média é %v\n", averageValue)
	fmt.Printf("O máximo é %v\n", maxValue)
}

func operation(operation string) (func(numeros ...int) float32, error) {
	switch operation {
	case "minimum":
		return minFunc, nil
	case "average":
		return averageFunc, nil
	case "maximum":
		return maxFunc, nil
	}
	return minFunc, errors.New("Operação não encontrada")
}

func minFunc(numeros ...int) float32 {
	var min int = numeros[0]
	for _, numero := range numeros {
		if numero < min {
			min = numero
		}
	}
	return float32(min)
}

func averageFunc(numeros ...int) float32 {
	var soma int
	for _, numero := range numeros {
		soma += numero
	}
	return float32(soma / len(numeros))
}

func maxFunc(numeros ...int) float32 {
	var max int = numeros[0]
	for _, numero := range numeros {
		if numero > max {
			max = numero
		}
	}
	return float32(max)
}
