package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var logger *log.Logger = log.Default()

func main() {
	rt := chi.NewRouter()

	rt.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("hello world get endpoint")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
