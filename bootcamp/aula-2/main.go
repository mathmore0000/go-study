package main

import (
	"fmt"

	"github.com/meli/teste/soma"
	"github.com/meli/teste/soma/sub"
)

func main() {
	fmt.Println(soma.Somar(5, 13))
	fmt.Println(sub.Subtrair(5, 13))
	// fmt.Println(soma.sub.Subtrair(5, 13)) // Erro!
}
