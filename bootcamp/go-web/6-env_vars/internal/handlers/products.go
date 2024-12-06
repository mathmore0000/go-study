package handlers

import (
	"encoding/json"
	"fmt"
	"main/internal/repository"
	"main/internal/service"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service *service.ProductService
}

type ProductRequest struct {
	repository.Product
}

type ProductResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message,omitempty"`
	Data    *repository.Product `json:"data,omitempty"`
}

type ProductsResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message,omitempty"`
	Data    *[]*repository.Product `json:"data"`
}

type ProductPatchRequest struct {
	Name        *string  `json:"name,omitempty"`
	Quantity    *int     `json:"quantity,omitempty"`
	CodeValue   *string  `json:"code_value,omitempty"`
	IsPublished *bool    `json:"is_published,omitempty"`
	Expiration  *string  `json:"expiration,omitempty"`
	Price       *float32 `json:"price,omitempty"`
}

var errDecodeJson ProductResponse = ProductResponse{
	Message: "Erro ao decodificar JSON", Status: http.StatusBadRequest,
}

var errProdNotFound ProductResponse = ProductResponse{
	Message: "Produto não encontrado", Status: http.StatusNotFound,
}

var errInvalidToken ProductResponse = ProductResponse{
	Message: "Token inválido", Status: http.StatusUnauthorized,
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(os.Getenv("TOKEN"))
	if os.Getenv("TOKEN") == "" || r.Header.Get("TOKEN") != os.Getenv("TOKEN") {
		json.NewEncoder(w).Encode(errInvalidToken)
		return

	}
	var productRequest ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		json.NewEncoder(w).Encode(errDecodeJson)
		return
	}

	var product *repository.Product = &repository.Product{
		Name:        productRequest.Name,
		Quantity:    productRequest.Quantity,
		CodeValue:   productRequest.CodeValue,
		IsPublished: productRequest.IsPublished,
		Expiration:  productRequest.Expiration,
		Price:       productRequest.Price,
	}

	if err := h.service.Create(product); err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	status := http.StatusCreated
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ProductResponse{Message: "Produto criado", Data: product, Status: status})
}

func (h *ProductHandler) PutProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if os.Getenv("TOKEN") == "" || r.Header.Get("TOKEN") != os.Getenv("TOKEN") {
		json.NewEncoder(w).Encode(errInvalidToken)
		return

	}
	var productRequest ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		json.NewEncoder(w).Encode(errDecodeJson)
		return
	}

	var product *repository.Product = &repository.Product{
		ID:          productRequest.ID,
		Name:        productRequest.Name,
		Quantity:    productRequest.Quantity,
		CodeValue:   productRequest.CodeValue,
		IsPublished: productRequest.IsPublished,
		Expiration:  productRequest.Expiration,
		Price:       productRequest.Price,
	}

	jaExistia, err := h.service.Put(product)
	if err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	if jaExistia {
		status := http.StatusOK
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(ProductResponse{Message: "Produto atualizado", Data: product, Status: status})
		return
	}
	status := http.StatusCreated
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ProductResponse{Message: "Produto criado", Data: product, Status: status})
}

func (h *ProductHandler) PatchProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if os.Getenv("TOKEN") == "" || r.Header.Get("TOKEN") != os.Getenv("TOKEN") {
		json.NewEncoder(w).Encode(errInvalidToken)
		return

	}

	// Obter o ID do produto a ser atualizado
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: "ID inválido", Status: http.StatusBadRequest})
		return
	}

	// Obter o produto existente
	product, ok := h.service.GetById(id)
	if !ok {
		json.NewEncoder(w).Encode(errProdNotFound)
		return
	}

	// Decodificar o corpo da requisição para ProductPatchRequest
	var productPatch ProductPatchRequest
	if err := json.NewDecoder(r.Body).Decode(&productPatch); err != nil {
		json.NewEncoder(w).Encode(errDecodeJson)
		return
	}

	// Atualizar apenas os campos fornecidos
	if productPatch.Name == nil {
		productPatch.Name = &product.Name
	}
	if productPatch.Quantity == nil {
		productPatch.Quantity = &product.Quantity
	}
	if productPatch.CodeValue == nil {
		productPatch.CodeValue = &product.CodeValue
	}
	if productPatch.IsPublished == nil {
		productPatch.IsPublished = &product.IsPublished
	}
	if productPatch.Expiration == nil {
		productPatch.Expiration = &product.Expiration
	}
	if productPatch.Price == nil {
		productPatch.Price = &product.Price
	}
	productToUpdate := &repository.Product{
		ID:          id,
		Name:        *productPatch.Name,
		Quantity:    *productPatch.Quantity,
		CodeValue:   *productPatch.CodeValue,
		IsPublished: *productPatch.IsPublished,
		Expiration:  *productPatch.Expiration,
		Price:       *productPatch.Price,
	}

	// Chamar o serviço para validar e atualizar o produto
	err = h.service.Patch(productToUpdate)
	if err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	status := http.StatusOK
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ProductResponse{Message: "Produto atualizado", Data: productToUpdate, Status: status})
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if os.Getenv("TOKEN") == "" || r.Header.Get("TOKEN") != os.Getenv("TOKEN") {
		json.NewEncoder(w).Encode(errInvalidToken)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: "id inválido", Status: http.StatusBadRequest})
		return
	}

	product, ok := h.service.GetById(id)
	if !ok {
		json.NewEncoder(w).Encode(errProdNotFound)
		return
	}

	err = h.service.Delete(product)

	if err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: err.Error(), Status: http.StatusNotFound})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) ListAllProducts(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetAll()

	w.Header().Set("Content-Type", "application/json")
	status := http.StatusCreated
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ProductsResponse{Data: products, Status: status})
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: "id inválido", Status: http.StatusBadRequest})
		return
	}

	product, ok := h.service.GetById(id)
	if !ok {
		json.NewEncoder(w).Encode(errProdNotFound)
		return
	}

	status := http.StatusCreated
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ProductResponse{Data: product, Status: status})
}

func (h *ProductHandler) GetProductSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var priceGtStr string = r.URL.Query().Get("priceGt")
	var priceGt float64
	var err error
	priceGt, err = strconv.ParseFloat(priceGtStr, 32)
	if priceGtStr != "" && err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: "priceGt inválido", Status: http.StatusBadRequest})
		return
	}

	products := h.service.GetProductsFiltered(float32(priceGt))

	status := http.StatusCreated
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ProductsResponse{Data: products, Status: status})
}
