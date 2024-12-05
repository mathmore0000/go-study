package main

import (
	"encoding/json"
	"log"
	"os"
)

type Aluno struct {
	Nome  string  `json:"nome"`
	Idade int     `json:"idade"`
	Nota  float64 `json:"nota"`
}

func main() {
	alunos := []Aluno{
		{"João", 28, 8.5},
		{"Maria", 22, 9},
		{"Pedro", 19, 7.5},
	}

	file, err := os.Create("alunos.json")
	if err != nil {
		log.Fatalf("Erro ao criar arquivo: %v", err)
	}

	encoder := json.NewEncoder(file)
	for _, aluno := range alunos {
		err := encoder.Encode(aluno)
		if err != nil {
			log.Fatalf("Erro ao codificar o aluno: %v", err)
		}
	}
}
