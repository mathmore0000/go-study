package handlers

import (
	"context"
	"main/internal/repository"
	"main/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestProductHandler_ListAllProducts(t *testing.T) {
	type fields struct {
		repo *repository.ProductRepository
	}
	// given
	// repository
	repoEmpty := &repository.ProductRepository{
		Products: make(map[int]*repository.Product, 1),
		NextID:   1,
	}
	repoFull := &repository.ProductRepository{
		Products: map[int]*repository.Product{
			1: {
				ID:          1,
				Name:        "Produto 1",
				Quantity:    5,
				CodeValue:   "2XP49S",
				IsPublished: true,
				Expiration:  "15/12/2024",
				Price:       99.99,
			}},
		NextID: 1,
	}
	tests := []struct {
		name       string
		fields     fields
		wantBody   string
		wantCode   int
		wantHeader http.Header
	}{
		{name: "Sucesso retorno vazio", fields: fields{repo: repoEmpty}, wantBody: `{"status":200,"data":[]}`, wantCode: http.StatusOK, wantHeader: http.Header{"Content-Type": []string{"application/json"}}},
		{name: "Sucesso retorno cheio", fields: fields{repo: repoFull}, wantBody: `
		{
			"status":200,
			"data":[{
				"code_value": "2XP49S",
				"expiration": "15/12/2024",
				"id": 1,
				"is_published": true,
				"name": "Produto 1",
				"price": 99.99,
				"quantity": 5
			}]
		}`,
			wantCode: http.StatusOK, wantHeader: http.Header{"Content-Type": []string{"application/json"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// service
			sv := service.NewProductService(tt.fields.repo)
			hd := NewProductHandler(sv)

			// when
			req := httptest.NewRequest("GET", "/products", nil)

			res := httptest.NewRecorder()
			hd.ListAllProducts(res, req)

			require.Equal(t, tt.wantCode, res.Code)
			require.JSONEq(t, tt.wantBody, res.Body.String())
			require.Equal(t, tt.wantHeader, res.Header())
		})
	}
}

func TestProductHandler_GetProductById(t *testing.T) {
	type fields struct {
		repo *repository.ProductRepository
	}
	type args struct {
		id string
	}
	// given
	// repository
	repoFull := &repository.ProductRepository{
		Products: map[int]*repository.Product{
			1: {
				ID:          1,
				Name:        "Produto 1",
				Quantity:    5,
				CodeValue:   "2XP49S",
				IsPublished: true,
				Expiration:  "15/12/2024",
				Price:       99.99,
			},
			2: {
				ID:          2,
				Name:        "Produto 2",
				Quantity:    9,
				CodeValue:   "4DFOS9S",
				IsPublished: false,
				Expiration:  "30/01/2025",
				Price:       999.99,
			}},
		NextID: 3,
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantBody   string
		wantCode   int
		wantHeader http.Header
	}{
		{name: "Sucesso retorno id inválido", args: args{id: "quarenta"}, wantBody: `{"status":400,"message":"id inválido"}`, wantCode: http.StatusBadRequest, wantHeader: http.Header{"Content-Type": []string{"application/json"}}},
		{name: "Sucesso retorno vazio", args: args{id: "3"}, wantBody: `{"status":404,"message":"Produto não encontrado"}`, wantCode: http.StatusNotFound, wantHeader: http.Header{"Content-Type": []string{"application/json"}}},
		{name: "Sucesso retorno válido", args: args{id: "1"}, wantBody: `
		{
			"status":200,
			"data":{
				"code_value": "2XP49S",
				"expiration": "15/12/2024",
				"id": 1,
				"is_published": true,
				"name": "Produto 1",
				"price": 99.99,
				"quantity": 5
			}
		}`,
			wantCode: http.StatusOK, wantHeader: http.Header{"Content-Type": []string{"application/json"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Inicializa o service e o handler
			sv := service.NewProductService(repoFull)
			hd := NewProductHandler(sv)

			// Cria a requisição com o ID do caso de teste
			req := httptest.NewRequest("GET", "/products/{id}", nil)
			res := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.args.id)

			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			// Serve a requisição usando o roteador
			hd.GetProductById(res, req)

			require.Equal(t, tt.wantCode, res.Code)
			require.JSONEq(t, tt.wantBody, res.Body.String())
			require.Equal(t, tt.wantHeader, res.Header())
		})
	}
}
