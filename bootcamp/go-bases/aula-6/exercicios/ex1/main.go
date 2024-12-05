package main

import (
	"fmt"
	"time"
)

/*
	Exercício 1 - Registro de alunos
	Uma universidade precisa registrar os alunos e gerar uma funcionalidade para imprimir os detalhes dos dados de cada aluno, como segue:

	Nome: [Primeiro nome do aluno]
	Sobrenome: [Sobrenome do aluno]
	ID: [ID do aluno]
	Data: [Data de admissão do aluno]

	Os valores entre colchetes devem ser substituídos pelos dados fornecidos pelos alunos.
	Para isso, é necessário gerar uma estrutura Students com as variáveis Name, Surname, DNI, Date e que tenha um método de detalhamento
*/

type Student struct {
	ID            int
	Name          string
	Surname       string
	AdmissionDate time.Time
}

func (student Student) print() {
	fmt.Printf("Nome: %v\nSobrenome: %v\nID: %v\nData: %v\n", student.Name, student.Surname, student.ID, student.AdmissionDate)
}

func main() {
	var student Student = Student{Name: "Matheus", Surname: "Moreira", ID: 1, AdmissionDate: time.Date(2024, 11, 18, 12, 0, 0, 0, time.Local)}
	student.print()
}
