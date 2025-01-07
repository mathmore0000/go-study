package internal

import "errors"

var (
	// ErrProductNotFound is an error that will be returned when a product is not found
	ErrProductNotFound = errors.New("repository: product not found")
	// ErrProductNotUnique is an error that will be returned when a product is not unique
	ErrProductNotUnique = errors.New("repository: product not unique")
	// ErrProductRelation is an error that will be returned when a product relation fails
	ErrProductRelation = errors.New("repository: product relation error")
)

// RepositoryProducts is an interface that represents a product repository
type RepositoryProducts interface {
	// GetOne returns a product by id
	GetOne(id int) (p Product, err error)
	// Store stores a product
	Store(p *Product) (err error)
	// Update updates a product
	Update(p *Product) (err error)
	// Delete deletes a product by id
	Delete(id int) (err error)

	// GetAll returns all products
	GetAll() (p []Product, err error)
}
