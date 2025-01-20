package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ProductMap(t *testing.T) {
	db := map[int]internal.Product{
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
	}
	r := repository.NewProductsMap(db)

	type Args struct {
		id internal.ProductQuery
	}
	tests := []struct {
		name string
		args Args
		want map[int]internal.Product
	}{
		{
			name: "Return all products by seller id",
			args: Args{
				id: internal.ProductQuery{
					Id: 1,
				},
			},
			want: map[int]internal.Product{
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
			},
		},
		{
			name: "Product not found",
			args: Args{
				id: internal.ProductQuery{
					Id: 3,
				},
			},
			want: map[int]internal.Product{},
		},
		{
			name: "Return all products",
			args: Args{
				id: internal.ProductQuery{},
			},
			want: map[int]internal.Product{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := r.SearchProducts(tt.args.id)
			require.Equal(t, tt.want, got)
		})
	}
}
