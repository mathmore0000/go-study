package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/*
Exercício 2 - Manipulação do Body

Vamos criar um endpoint chamado /greetings.

	Com uma pequena estrutura com nome e sobrenome que, quando colada, deve responder no texto "Hello + nome + sobrenome".

O ponto de extremidade deve ser um método POST
O package JSON deve ser usado para resolver o exercício.
A resposta deve seguir a seguinte estrutura: "Hello Andrea Rivas".
A estrutura deve ter a seguinte aparência:

	{
		"firstName": "Andrea",
		"lastName": "Rivas".
	}
*/

type GreetinsRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	rt := chi.NewRouter()
	rt.Use(middleware.Logger)

	rt.Get("/greetings", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		var greetingsRequset GreetinsRequest
		err = json.Unmarshal(body, &greetingsRequset)
		if err != nil {
			http.Error(w, "Erro ao decodificar json", http.StatusBadRequest)
			return
		}
		if greetingsRequset.FirstName == "" || greetingsRequset.LastName == "" {
			http.Error(w, "firstName e lastName são obrigatórios", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello, %s %s", greetingsRequset.FirstName, greetingsRequset.LastName)))
	})

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
