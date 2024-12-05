package handlers

import (
	"encoding/json"
	"main/internal/repository"
	"main/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product repository.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) ListAllProducts(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetAll()

	productsMarsharled, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Erro ao deserializar produtos", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(productsMarsharled))
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	product, ok := h.service.GetById(id)
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
}

func (h *ProductHandler) GetProductSearch(w http.ResponseWriter, r *http.Request) {
	var priceGtStr string = r.URL.Query().Get("priceGt")
	var priceGt float64
	var err error
	priceGt, err = strconv.ParseFloat(priceGtStr, 32)
	if priceGtStr != "" && err != nil {
		http.Error(w, "Preço inválido", http.StatusBadRequest)
		return
	}

	product := h.service.GetProductsFiltered(float32(priceGt))

	productsMarsharled, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Erro ao deserializar produtos", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(productsMarsharled))
}

// func initializeProducts() {
// 	// load json
// 	ProductsById = make(map[int]*Product, len(products))
// 	readProductsFromJson("products.json", &products)

// 	for _, product := range products {
// 		ProductsById[product.ID] = product
// 	}
// }

// func readProductsFromJson(jsonPath string, values *[]*Product) {
// 	// Open the JSON file
// 	jsonFile, err := os.Open(jsonPath)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer jsonFile.Close()

// 	// Read the file's content into a byte slice
// 	byteValue, err := io.ReadAll(jsonFile)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// Unmarshal the JSON data
// 	err = json.Unmarshal(byteValue, values)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
// 		os.Exit(1)
// 	}
// }
