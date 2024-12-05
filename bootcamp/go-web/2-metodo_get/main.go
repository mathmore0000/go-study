package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float32 `json:"price"`
}

var products []*Product = []*Product{}
var ProductsById map[int]*Product

func main() {
	initializeProducts()
	rt := chi.NewRouter()
	rt.Use(middleware.Logger)

	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	rt.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		productsMarsharled, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "Erro ao deserializar produtos", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(productsMarsharled))
	})

	rt.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		product, ok := ProductsById[id]
		if !ok {
			http.Error(w, "Produto não encontrado", http.StatusNotFound)
			return
		}

		productsMarsharled, err := json.Marshal(product)
		if err != nil {
			http.Error(w, "Erro ao deserializar produtos", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(productsMarsharled))
	})
	rt.Get("/products/search", func(w http.ResponseWriter, r *http.Request) {
		var priceGtStr string = r.URL.Query().Get("priceGt")
		var priceGt float64
		var err error
		priceGt, err = strconv.ParseFloat(priceGtStr, 32)
		if priceGtStr != "" && err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		product := getProducts(float32(priceGt))

		productsMarsharled, err := json.Marshal(product)
		if err != nil {
			http.Error(w, "Erro ao deserializar produtos", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(productsMarsharled))
	})

	if err := http.ListenAndServe(":80", rt); err != nil {
		panic(err)
	}
}

func getProducts(priceGt float32) []*Product {
	var ps []*Product = []*Product{}

	for _, product := range products {
		if product.Price > priceGt {
			ps = append(ps, product)
		}
	}
	return ps
}

func initializeProducts() {
	// load json
	ProductsById = make(map[int]*Product, len(products))
	readProductsFromJson("products.json", &products)

	for _, product := range products {
		ProductsById[product.ID] = product
	}
}

func readProductsFromJson(jsonPath string, values *[]*Product) {
	// Open the JSON file
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	// Read the file's content into a byte slice
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal the JSON data
	err = json.Unmarshal(byteValue, values)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
}
