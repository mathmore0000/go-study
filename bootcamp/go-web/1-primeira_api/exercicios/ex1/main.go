package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
Exercício 1 - Teste de ping
Vamos criar um aplicativo da Web com o package net/http nativo do Go, que tem um endpoint /ping que, quando colado, responde com um texto que diz "pong".

O endpoint deve ser um método GET
A resposta do pong deve ser enviada como texto, NÃO como JSON.
*/
func main() {
	rt := chi.NewRouter()

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
