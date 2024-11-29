package main

import (
	"fmt"
	"time"
)

/*
	Exercício 2 - Employee
	Uma empresa precisa fazer um bom gerenciamento de seus funcionários.
		Para isso, criaremos um pequeno programa que nos ajudará a gerenciar corretamente esses funcionários. Os objetivos são:

	Criar uma estrutura Person com os campos ID, Name, DateOfBirth.
	Criar uma estrutura Employee com os campos: ID, Position e uma composição com a estrutura Person.
	Criar um método para a estrutura Employee chamado PrintEmployee(), que imprimirá os campos de um funcionário.
	Instanciar na função main() uma Person e um Employee carregando seus respectivos campos e, finalmente, executar o método PrintEmployee().
	Se você conseguir realizar esse pequeno programa, poderá ajudar a empresa a resolver o problema de gerenciamento dos funcionários.
*/

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func (employee *Employee) PrintEmployee() {
	fmt.Printf("ID -> %v | Position -> %v | Person.ID -> %v | Person.Name -> %v | Person.DateOfBirth -> %v\n", employee.ID, employee.Position, employee.Person.ID, employee.Person.Name, employee.Person.DateOfBirth)
}

func main() {
	var person Person = Person{ID: 1, Name: "Daniel Filho", DateOfBirth: time.Date(2000, 9, 2, 12, 0, 0, 0, time.Local)}
	var employee Employee = Employee{ID: 1, Position: "Software Developer", Person: person}

	employee.PrintEmployee()

}
