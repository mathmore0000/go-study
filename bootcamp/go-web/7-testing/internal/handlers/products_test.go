package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"main/internal/repository"
	"main/internal/service"
	"net/http"
	"net/http/httptest"
	"os"
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

			// then
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
		{name: "Falha retorno id inválido", args: args{id: "quarenta"}, wantBody: `{"status":400,"message":"id inválido"}`, wantCode: http.StatusBadRequest, wantHeader: http.Header{"Content-Type": []string{"application/json"}}},
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

func TestProductHandler_GetProductSearch(t *testing.T) {
	type fields struct {
		service *service.ProductService
	}
	type args struct {
		priceGt string
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
		{name: "Falha retorno priceGt inválido", args: args{priceGt: "duzentos"}, wantBody: `{"status":400,"message":"priceGt inválido"}`, wantCode: http.StatusBadRequest, wantHeader: http.Header{"Content-Type": []string{"application/json"}}},
		{name: "Sucesso retorno vazio", args: args{priceGt: "1000"}, wantBody: `{"status":200,"data":null}`, wantCode: http.StatusOK, wantHeader: http.Header{"Content-Type": []string{"application/json"}}},
		{name: "Sucesso retorno válido", args: args{priceGt: "900"}, wantBody: `
		{
			"status":200,
			"data":[{
				"code_value": "4DFOS9S",
				"expiration": "30/01/2025",
				"id": 2,
				"is_published": false,
				"name": "Produto 2",
				"price": 999.99,
				"quantity": 9
			}]
		}`,
			wantCode: http.StatusOK, wantHeader: http.Header{"Content-Type": []string{"application/json"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Inicializa o service e o handler
			sv := service.NewProductService(repoFull)
			hd := NewProductHandler(sv)

			// Cria a requisição com o ID do caso de teste
			req := httptest.NewRequest("GET", fmt.Sprintf("/products/search?priceGt=%s", tt.args.priceGt), nil)
			res := httptest.NewRecorder()

			// Serve a requisição usando o roteador
			hd.GetProductSearch(res, req)

			require.Equal(t, tt.wantCode, res.Code)
			require.JSONEq(t, tt.wantBody, res.Body.String())
			require.Equal(t, tt.wantHeader, res.Header())
		})
	}
}

func TestProductHandler_CreateProduct(t *testing.T) {
	type fields struct {
		service *service.ProductService
	}
	type args struct {
		newProduct repository.Product
		token      string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantBody   string
		wantCode   int
		wantHeader http.Header
	}{
		{
			name: "Falha falta de token",
			args: args{
				token:      "secret-token-errado",
				newProduct: repository.Product{Name: "Produto 1", Quantity: 5, CodeValue: "2XP49S", IsPublished: true, Expiration: "15/12/2024", Price: 99.99}},
			wantBody:   `{"status":401,"message":"Token inválido"}`,
			wantCode:   http.StatusUnauthorized,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Sucesso na criação do produto",
			args: args{
				token:      "secret-token",
				newProduct: repository.Product{Name: "Produto 1", Quantity: 5, CodeValue: "2XP49S", IsPublished: true, Expiration: "15/12/2024", Price: 99.99}},
			wantBody:   `{"status":201,"message":"Produto criado","data":{"code_value":"2XP49S", "expiration":"15/12/2024", "id":1, "is_published":true, "name":"Produto 1", "price":99.99, "quantity":5}}`,
			wantCode:   http.StatusCreated,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
	}
	os.Setenv("TOKEN", "secret-token")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// given
			// Inicializa o service e o handler
			sv := service.NewProductService(&repository.ProductRepository{
				Products: make(map[int]*repository.Product, 1),
				NextID:   1,
				JsonPath: "../mocks/products.json",
			})

			// when
			marshaledNewProduct, _ := json.Marshal(tt.args.newProduct)
			req := httptest.NewRequest("POST", "/products", bytes.NewReader(marshaledNewProduct))
			req.Header.Set("TOKEN", tt.args.token)
			res := httptest.NewRecorder()

			hd := NewProductHandler(sv)

			hd.CreateProduct(res, req)

			// then
			require.Equal(t, tt.wantCode, res.Code)
			require.JSONEq(t, tt.wantBody, res.Body.String())
			require.Equal(t, tt.wantHeader, res.Header())
		})
	}
}

func TestProductHandler_PutProduct(t *testing.T) {
	type args struct {
		token   string
		product repository.Product
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
			},
		},
		NextID:   2,
		JsonPath: "../mocks/products.json",
	}
	tests := []struct {
		name       string
		args       args
		wantBody   string
		wantCode   int
		wantHeader http.Header
	}{
		{
			name: "Falha falta de token",
			args: args{
				token: "token-invalido",
				product: repository.Product{
					ID:          1,
					Name:        "Produto Atualizado",
					Quantity:    10,
					CodeValue:   "2XP49S",
					IsPublished: true,
					Expiration:  "15/12/2025",
					Price:       199.99,
				},
			},
			wantBody:   `{"status":401,"message":"Token inválido"}`,
			wantCode:   http.StatusUnauthorized,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Sucesso na atualização do produto existente",
			args: args{
				token: "secret-token",
				product: repository.Product{
					ID:          1,
					Name:        "Produto Atualizado",
					Quantity:    10,
					CodeValue:   "2XP49S",
					IsPublished: true,
					Expiration:  "15/12/2025",
					Price:       199.99,
				},
			},
			wantBody:   `{"status":200,"message":"Produto atualizado","data":{"id":1,"name":"Produto Atualizado","quantity":10,"code_value":"2XP49S","is_published":true,"expiration":"15/12/2025","price":199.99}}`,
			wantCode:   http.StatusOK,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Sucesso na criação de novo produto",
			args: args{
				token: "secret-token",
				product: repository.Product{
					ID:          2,
					Name:        "Novo Produto",
					Quantity:    15,
					CodeValue:   "NEWCODE",
					IsPublished: false,
					Expiration:  "01/01/2026",
					Price:       299.99,
				},
			},
			wantBody:   `{"status":201,"message":"Produto criado","data":{"id":2,"name":"Novo Produto","quantity":15,"code_value":"NEWCODE","is_published":false,"expiration":"01/01/2026","price":299.99}}`,
			wantCode:   http.StatusCreated,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
	}
	os.Setenv("TOKEN", "secret-token")
	defer os.Unsetenv("TOKEN")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Inicializa o service e o handler com o repositório
			sv := service.NewProductService(repoFull)
			hd := NewProductHandler(sv)

			// Serializa o produto para JSON
			marshaledProduct, _ := json.Marshal(tt.args.product)

			// Cria a requisição
			req := httptest.NewRequest("PUT", "/products", bytes.NewReader(marshaledProduct))
			req.Header.Set("TOKEN", tt.args.token)
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()

			// Chama o handler
			hd.PutProduct(res, req)

			// Verifica o resultado
			require.Equal(t, tt.wantCode, res.Code)
			require.JSONEq(t, tt.wantBody, res.Body.String())
			require.Equal(t, tt.wantHeader.Get("Content-Type"), res.Header().Get("Content-Type"))
		})
	}
}

func TestProductHandler_PatchProduct(t *testing.T) {
	type args struct {
		token        string
		id           string
		productPatch ProductPatchRequest
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
			},
		},
		NextID:   2,
		JsonPath: "../mocks/products.json",
	}
	tests := []struct {
		name       string
		args       args
		wantBody   string
		wantCode   int
		wantHeader http.Header
	}{
		{
			name: "Falha falta de token",
			args: args{
				token: "token-invalido",
				id:    "1",
				productPatch: ProductPatchRequest{
					Name: ptrString("Produto Patch"),
				},
			},
			wantBody:   `{"status":401,"message":"Token inválido"}`,
			wantCode:   http.StatusUnauthorized,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Falha ID inválido",
			args: args{
				token: "secret-token",
				id:    "abc",
				productPatch: ProductPatchRequest{
					Name: ptrString("Produto Patch"),
				},
			},
			wantBody:   `{"status":400,"message":"id inválido"}`,
			wantCode:   http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Falha produto não encontrado",
			args: args{
				token: "secret-token",
				id:    "2",
				productPatch: ProductPatchRequest{
					Name: ptrString("Produto Patch"),
				},
			},
			wantBody:   `{"status":404,"message":"Produto não encontrado"}`,
			wantCode:   http.StatusNotFound,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Sucesso na atualização parcial do produto",
			args: args{
				token: "secret-token",
				id:    "1",
				productPatch: ProductPatchRequest{
					Name:     ptrString("Produto Atualizado via Patch"),
					Quantity: ptrInt(20),
				},
			},
			wantBody:   `{"status":200,"message":"Produto atualizado","data":{"id":1,"name":"Produto Atualizado via Patch","quantity":20,"code_value":"2XP49S","is_published":true,"expiration":"15/12/2024","price":99.99}}`,
			wantCode:   http.StatusOK,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
	}
	os.Setenv("TOKEN", "secret-token")
	defer os.Unsetenv("TOKEN")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Inicializa o service e o handler com o repositório
			sv := service.NewProductService(repoFull)
			hd := NewProductHandler(sv)

			// Serializa o productPatch para JSON
			marshaledProductPatch, _ := json.Marshal(tt.args.productPatch)

			// Cria a requisição com o ID especificado
			req := httptest.NewRequest("PATCH", fmt.Sprintf("/products/%s", tt.args.id), bytes.NewReader(marshaledProductPatch))
			req.Header.Set("TOKEN", tt.args.token)
			req.Header.Set("Content-Type", "application/json")

			// Configura o RouteContext para o ID
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.args.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			res := httptest.NewRecorder()

			// Chama o handler
			hd.PatchProduct(res, req)

			// Verifica o resultado
			require.Equal(t, tt.wantCode, res.Code)
			require.JSONEq(t, tt.wantBody, res.Body.String())
			require.Equal(t, tt.wantHeader.Get("Content-Type"), res.Header().Get("Content-Type"))
		})
	}
}

func TestProductHandler_DeleteProduct(t *testing.T) {
	type args struct {
		token string
		id    string
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
			},
		},
		NextID:   2,
		JsonPath: "../mocks/products.json",
	}
	tests := []struct {
		name       string
		args       args
		wantBody   string
		wantCode   int
		wantHeader http.Header
	}{
		{
			name: "Falha falta de token",
			args: args{
				token: "token-invalido",
				id:    "1",
			},
			wantBody:   `{"status":401,"message":"Token inválido"}`,
			wantCode:   http.StatusUnauthorized,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Falha ID inválido",
			args: args{
				token: "secret-token",
				id:    "abc",
			},
			wantBody:   `{"status":400,"message":"id inválido"}`,
			wantCode:   http.StatusBadRequest,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Falha produto não encontrado",
			args: args{
				token: "secret-token",
				id:    "2",
			},
			wantBody:   `{"status":404,"message":"Produto não encontrado"}`,
			wantCode:   http.StatusNotFound,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
		{
			name: "Sucesso na exclusão do produto",
			args: args{
				token: "secret-token",
				id:    "1",
			},
			wantBody:   `{"status":200,"message":"Produto deletado"}`,
			wantCode:   http.StatusOK,
			wantHeader: http.Header{"Content-Type": []string{"application/json"}},
		},
	}
	os.Setenv("TOKEN", "secret-token")
	defer os.Unsetenv("TOKEN")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Inicializa o service e o handler com o repositório
			sv := service.NewProductService(repoFull)
			hd := NewProductHandler(sv)

			// Cria a requisição com o ID especificado
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/products/%s", tt.args.id), nil)
			req.Header.Set("TOKEN", tt.args.token)
			res := httptest.NewRecorder()

			// Configura o RouteContext para o ID
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.args.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			// Chama o handler
			hd.DeleteProduct(res, req)

			// Verifica o resultado
			require.Equal(t, tt.wantCode, res.Code)
			require.JSONEq(t, tt.wantBody, res.Body.String())
			// No caso do DELETE bem-sucedido, não há corpo na resposta
			require.Equal(t, tt.wantHeader.Get("Content-Type"), res.Header().Get("Content-Type"))
		})
	}
}

func ptrString(s string) *string { return &s }
func ptrInt(i int) *int          { return &i }
