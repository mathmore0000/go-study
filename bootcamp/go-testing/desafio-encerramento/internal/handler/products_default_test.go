package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var url_base = "/products"

type RepositoryProductsMock struct {
	mock.Mock
}

func (m *RepositoryProductsMock) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	args := m.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}

func Test_Get(t *testing.T) {
	type TestCases struct {
		description       string
		id                any
		expectedCode      int
		expectedBody      string
		mock              func() *RepositoryProductsMock
		expectedMockCalls int
	}

	testCases := []TestCases{
		{
			description:  "case 1 - success: Return all products by id 1",
			id:           1,
			expectedCode: 200,
			expectedBody: `
				{
					"message": "success",
					"data": {
						"1": {
							"id": 1,
							"description": "product 1",
							"price": 100,
							"seller_id": 1
						}
					}
				}
			`,
			mock: func() *RepositoryProductsMock {
				mk := new(RepositoryProductsMock)
				mk.On("SearchProducts", internal.ProductQuery{Id: 1}).Return(map[int]internal.Product{
					1: {
						Id: 1,
						ProductAttributes: internal.ProductAttributes{
							Description: "product 1",
							Price:       100,
							SellerId:    1,
						},
					},
				}, nil)
				return mk
			},
			expectedMockCalls: 1,
		},
		{
			description:  "case 2 - success: Return all products",
			id:           0,
			expectedCode: 200,
			expectedBody: `{"data":{"1":{"id":1,"description":"product 1","price":100,"seller_id":1},"2":{"id":2,"description":"product 2","price":200,"seller_id":1},"3":{"id":3,"description":"product 3","price":300,"seller_id":2}},"message":"success"}`,
			mock: func() *RepositoryProductsMock {
				mk := new(RepositoryProductsMock)
				mk.On("SearchProducts", internal.ProductQuery{}).Return(map[int]internal.Product{
					1: {
						Id: 1,
						ProductAttributes: internal.ProductAttributes{
							Description: "product 1",
							Price:       100,
							SellerId:    1,
						},
					},
					2: {
						Id: 2,
						ProductAttributes: internal.ProductAttributes{
							Description: "product 2",
							Price:       200,
							SellerId:    1,
						},
					},
					3: {
						Id: 3,
						ProductAttributes: internal.ProductAttributes{
							Description: "product 3",
							Price:       300,
							SellerId:    2,
						},
					},
				}, nil)
				return mk
			},
			expectedMockCalls: 1,
		},
		{
			description:  "case 3 - error: id is not a number",
			id:           "s",
			expectedCode: 400,
			expectedBody: `{
				"status":"Bad Request",
				"message":"invalid id"
			}`,
			mock: func() *RepositoryProductsMock {
				mk := new(RepositoryProductsMock)
				mk.On("SearchProducts", internal.ProductQuery{}).Return(map[int]internal.Product{}, nil)
				return mk
			},
			expectedMockCalls: 0,
		},
		{
			description:  "case 4 - error: internal error",
			id:           0,
			expectedCode: 500,
			expectedBody: `{
				"status":"Internal Server Error",
				"message":"internal error"
			}`,
			mock: func() *RepositoryProductsMock {
				mk := new(RepositoryProductsMock)
				mk.On("SearchProducts", internal.ProductQuery{}).Return(map[int]internal.Product{}, errors.New("internal error"))
				return mk
			},
			expectedMockCalls: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// handler and its dependencies
			rp := tc.mock()
			hd := handler.NewProductsDefault(rp)
			hdFunc := hd.Get()

			// http request and response
			id := fmt.Sprintf("?id=%v", tc.id)
			request := httptest.NewRequest("GET", url_base+id, nil)
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()

			// WHEN
			hdFunc(response, request)

			// THEN
			require.Equal(t, tc.expectedCode, response.Code)
			require.JSONEq(t, tc.expectedBody, response.Body.String())
			rp.AssertNumberOfCalls(t, "SearchProducts", tc.expectedMockCalls)
		})
	}
}
