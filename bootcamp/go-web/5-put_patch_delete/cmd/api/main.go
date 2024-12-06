package main

import (
	"main/internal/handlers"
	"main/internal/repository"
	"main/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	rt := chi.NewRouter()
	rt.Use(middleware.Logger)

	// repository
	db := repository.NewProductRepository("../products.json")

	// service
	productService := service.NewProductService(db)

	// controller
	productHandler := handlers.NewProductHandler(productService)

	rt.Route("/products", func(r chi.Router) {
		r.Post("/", productHandler.CreateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Put("/", productHandler.PutProduct)
		r.Patch("/{id}", productHandler.PatchProduct)
		r.Get("/", productHandler.ListAllProducts)
		r.Get("/{id}", productHandler.GetProductById)
		r.Get("/search", productHandler.GetProductSearch)
	})

	if err := http.ListenAndServe(":80", rt); err != nil {
		panic(err)
	}
}
