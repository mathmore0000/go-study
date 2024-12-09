package main

import (
	"encoding/json"
	"log"
	"main/internal/handlers"
	"main/internal/repository"
	"main/internal/service"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	rt := chi.NewRouter()
	// rt.Use(middleware.Logger)

	// repository
	db := repository.NewProductRepository("../../internal/repository/products.json")

	// service
	productService := service.NewProductService(db)

	// controller
	productHandler := handlers.NewProductHandler(productService)

	rt.Use(loggingMiddleware)
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", productHandler.ListAllProducts)
		r.Get("/search", productHandler.GetProductSearch)
		r.Get("/{id}", productHandler.GetProductById)

		r.Group(func(r chi.Router) {
			r.Use(authMiddleware)

			r.Post("/", productHandler.CreateProduct)
			r.Put("/", productHandler.PutProduct)
			r.Patch("/{id}", productHandler.PatchProduct)
			r.Delete("/{id}", productHandler.DeleteProduct)
		})
	})

	if err := http.ListenAndServe(":80", rt); err != nil {
		panic(err)
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("TOKEN") != os.Getenv("TOKEN") {
			status := http.StatusUnauthorized
			w.WriteHeader(status)
			json.NewEncoder(w).Encode("Token inválido")
			return
		}
		next.ServeHTTP(w, r) // Chama o próximo handler na cadeia
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	bytesWritten int
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture start time
		startTime := time.Now()

		// Create a custom ResponseWriter to capture the response size
		lrw := &loggingResponseWriter{ResponseWriter: w}

		// Call the next handler
		next.ServeHTTP(lrw, r)

		// Log the required information
		method := r.Method
		url := r.URL.String()
		timestamp := startTime.Format(time.RFC3339)
		bytesWritten := lrw.bytesWritten

		log.Printf("Método: %s | Data e Hora: %s | URL: %s | Tamanho da Resposta: %d bytes",
			method, timestamp, url, bytesWritten)
	})
}
