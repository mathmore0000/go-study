package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

type ProductRepository struct {
	products map[int]*Product
	nextID   int
}

func NewProductRepository(jsonPath string) *ProductRepository {
	repo := &ProductRepository{
		products: make(map[int]*Product),
		nextID:   1,
	}
	repo.loadProducts(jsonPath)
	return repo
}

func (r *ProductRepository) GetById(id int) (*Product, bool) {
	product, ok := r.products[id]
	return product, ok
}

func (r *ProductRepository) GetAll() *[]*Product {
	products := make([]*Product, 0, len(r.products))
	for _, p := range r.products {
		products = append(products, p)
	}
	return &products
}

func (r *ProductRepository) SearchByPrice(priceGt float32) *[]*Product {
	var products []*Product
	for _, p := range r.products {
		if p.Price > priceGt {
			products = append(products, p)
		}
	}
	return &products
}

func (r *ProductRepository) Create(product *Product) error {
	// Verificar se o code_value já existe
	for _, p := range r.products {
		if p.CodeValue == product.CodeValue {
			return fmt.Errorf("Product with code_value '%s' already exists", product.CodeValue)
		}
	}

	product.ID = r.nextID
	r.nextID++
	r.products[product.ID] = product
	return nil
}

func (r *ProductRepository) loadProducts(jsonPath string) {
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

	var products []*Product
	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	for _, product := range products {
		r.products[product.ID] = product
	}
	r.nextID = len(r.products) + 1
}
