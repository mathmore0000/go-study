package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("A")
	go task("B")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i, "ANONIMOUS")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(15 * time.Second)
}
