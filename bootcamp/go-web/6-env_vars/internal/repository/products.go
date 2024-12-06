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
	jsonPath string
}

func NewProductRepository(jsonPath string) *ProductRepository {
	repo := &ProductRepository{
		products: make(map[int]*Product),
		nextID:   1,
		jsonPath: jsonPath,
	}
	repo.loadProducts()
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

func (r *ProductRepository) Delete(product *Product) error {
	delete(r.products, product.ID)
	err := r.writeToProductsToJson()
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) ExistsByCodeValue(product *Product) bool {
	for _, p := range r.products {
		if p.ID != product.ID && p.CodeValue == product.CodeValue {
			return true
		}
	}
	return false
}

func (r *ProductRepository) Create(product *Product) error {
	// Verificar se o code_value já existe
	product.ID = r.nextID
	r.nextID++
	r.products[product.ID] = product

	err := r.writeToProductsToJson()
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Put(product *Product) (alreadyExisted bool, err error) {
	// Verificar se o code_value já existe
	_, ok := r.products[product.ID]

	if ok {
		alreadyExisted = true
		r.products[product.ID] = product
	} else {
		product.ID = r.nextID
		r.nextID++
		r.products[product.ID] = product
	}
	err = r.writeToProductsToJson()
	if err != nil {
		return false, err
	}
	return
}

func (r *ProductRepository) Update(product *Product) error {
	r.products[product.ID] = product
	err := r.writeToProductsToJson()
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Patch(product *Product) error {
	r.products[product.ID] = product

	return nil
}

// func (r *ProductRepository) Save(product *Product) error {}

func (r *ProductRepository) writeToProductsToJson() error {
	// Abrir o arquivo JSON para escrita (criando ou truncando o arquivo)
	jsonFile, err := os.Create(r.jsonPath)
	if err != nil {
		return fmt.Errorf("Erro ao abrir arquivo para escrita: %v", err)
	}
	defer jsonFile.Close()

	// Codificar os produtos em JSON e escrever diretamente no arquivo
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ") // Opcional: para formatar o JSON de forma legível
	err = encoder.Encode(r.products)
	if err != nil {
		return fmt.Errorf("Erro ao codificar JSON: %v", err)
	}

	return nil
}

func (r *ProductRepository) loadProducts() {
	// Open the JSON file
	jsonFile, err := os.Open(r.jsonPath)
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

	var products map[int]*Product
	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	r.nextID = len(r.products) + 1
	r.products = products
}
