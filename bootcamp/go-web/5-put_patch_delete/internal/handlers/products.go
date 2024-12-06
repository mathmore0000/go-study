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

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var productRequest ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: "Erro ao decodificar JSON", Status: http.StatusBadRequest})
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

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(ProductResponse{Message: "id inválido", Status: http.StatusBadRequest})
		return
	}

	product, ok := h.service.GetById(id)
	if !ok {
		json.NewEncoder(w).Encode(ProductResponse{Message: "Produto não encontrado", Status: http.StatusNotFound})
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
		json.NewEncoder(w).Encode(ProductResponse{Message: "Produto não encontrado", Status: http.StatusNotFound})
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
