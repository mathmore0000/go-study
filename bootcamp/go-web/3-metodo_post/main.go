package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
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

type ProductBodyRequest struct {
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

	rt.Post("/products", func(w http.ResponseWriter, r *http.Request) {
		// receber o body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		// validar request body
		newProduct := &Product{}
		err = validateProductBodyRequest(body, newProduct)
		if err != nil {
			fmt.Println("err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// inserir
		newProduct.ID = len(products) + 1
		ProductsById[newProduct.ID] = newProduct
		products = append(products, newProduct)

		// if err return err
		newProductMarsharled, err := json.Marshal(newProduct)
		if err != nil {
			http.Error(w, "Erro ao deserializar produtos", http.StatusInternalServerError)
			return
		}

		// return success
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(newProductMarsharled))
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

		product := getProductsFiltered(float32(priceGt))

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

func getProductsFiltered(priceGt float32) []*Product {
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

/*
Nenhum dado pode estar vazio, exceto is_published (vazio indica um valor false).
O campo code_value deve ser exclusivo para cada produto.
Os tipos de dados devem corresponder aos definidos no enunciado do problema.
A data de validade deve ter o formato: XX/XX/XXXXXX, e também devemos verificar se o dia, o mês e o ano são valores válidos.
*/
func validateProductBodyRequest(body []byte, product *Product) error {
	// Deserializar o corpo da requisição para o struct Product
	fmt.Println(body)
	err := json.Unmarshal(body, product)
	if err != nil {
		fmt.Errorf("Erro ao decodificar JSON: %v", err)
	}

	// Verificar se os campos obrigatórios não estão vazios
	if product.Name == "" {
		return errors.New("O campo 'name' não pode ser vazio")
	}
	if product.Quantity == 0 {
		return errors.New("O campo 'quantity' não pode ser zero")
	}
	if product.CodeValue == "" {
		return errors.New("O campo 'code_value' não pode ser vazio")
	}
	if product.Expiration == "" {
		return errors.New("O campo 'expiration' não pode ser vazio")
	}
	if product.Price == 0 {
		return errors.New("O campo 'price' não pode ser zero")
	}
	// 'is_published' pode ser vazio (valor padrão false)

	// Verificar se 'code_value' é único
	for _, p := range products {
		if p.CodeValue == product.CodeValue {
			return errors.New("O 'code_value' já existe")
		}
	}

	// Validar o formato da data de validade (XX/XX/XXXXXX)
	// e verificar se dia, mês e ano são valores válidos
	dateRegex := regexp.MustCompile(`^(\d{2})/(\d{2})/(\d{4})$`)
	matches := dateRegex.FindStringSubmatch(product.Expiration)
	if matches == nil {
		return errors.New("O campo 'expiration' deve estar no formato DD/MM/AAAA")
	}

	dayStr, monthStr, yearStr := matches[1], matches[2], matches[3]

	// Converter strings para inteiros
	day, err := strconv.Atoi(dayStr)
	if err != nil || day < 1 || day > 31 {
		return errors.New("Dia inválido na data de validade")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		return errors.New("Mês inválido na data de validade")
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil || year < 1 {
		return errors.New("Ano inválido na data de validade")
	}

	// Opcional: Poderíamos verificar se a data é válida considerando meses com 30/31 dias e anos bissextos

	return nil
}
