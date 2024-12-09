package main

import (
	"main/internal/handlers"
	"main/internal/repository"
	"main/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	rt := chi.NewRouter()
	rt.Use(middleware.Logger)

	// repository
	db := repository.NewProductRepository("../../internal/repository/products.json")

	// service
	productService := service.NewProductService(db)

	// controller
	productHandler := handlers.NewProductHandler(productService)

	rt.Route("/products", func(r chi.Router) {
		r.Get("/", productHandler.ListAllProducts)
		r.Get("/search", productHandler.GetProductSearch)
		r.Get("/{id}", productHandler.GetProductById)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/", productHandler.PutProduct)
		r.Patch("/{id}", productHandler.PatchProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	if err := http.ListenAndServe(":80", rt); err != nil {
		panic(err)
	}
}
