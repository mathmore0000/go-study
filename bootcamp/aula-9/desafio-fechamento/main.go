package main

import (
	"fmt"
	"time"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

/*
	Criar um programa que sirva como ferramenta para calcular diferentes dados estatísticos.
	Para isso, você deve clonar este repositório que contém um arquivo .csv com dados gerados e um esqueleto do projeto.

	Exercício 1:
		Uma função que calcula quantas pessoas viajam para um determinado país:
			func GetTotalTickets(destination string) (int, error) {}

	Exercício 2:
		Uma ou mais funções que calculam quantas pessoas viajam no início da manhã (0 → 6), manhã (7 → 12), tarde (13 → 19) e noite (20 → 23):

	Exercício 3:
		Calcule a porcentagem de pessoas que viajam para um determinado país em um determinado dia, em relação ao restante:

	Exercício 4:
		Crie testes de unidade para cada um dos requisitos acima, com um mínimo de 2 casos por requisito:
*/

var csvFilePath string = "tickets.csv"

func main() {
	tickets.InitializeTickets(csvFilePath)
	destination := "Brazil"
	totalTicketsByDestination, err := tickets.GetTotalTicketsByDestination(destination)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Total number of tickets in %v -> %v\n", destination, totalTicketsByDestination)
	}

	totalTicketsByEvening, timeOfDay := tickets.GetTotalTicketsByTime(time.Date(2024, 12, 02, 1, 0, 0, 0, time.Local))
	fmt.Printf("Total number of tickets in the %v -> %v\n", timeOfDay, totalTicketsByEvening)

	percentageOfTicketsByDestination, err := tickets.GetPercentageByDestination(destination)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The choice for %v as a destination represents %v%% of the tickets\n", destination, percentageOfTicketsByDestination)
	}
}
