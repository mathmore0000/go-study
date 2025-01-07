package internal

import "errors"

var (
	// ErrWarehouseNotFound is an error that will be returned when a warehouse is not found
	ErrWarehouseNotFound = errors.New("repository: warehouse not found")
	// ErrWarehouseNotUnique is an error that will be returned when a warehouse is not unique
	ErrWarehouseNotUnique = errors.New("repository: warehouse not unique")
	// ErrWarehouseRelation is an error that will be returned when a warehouse relation fails
	ErrWarehouseRelation = errors.New("repository: warehouse relation error")
)

// RepositoryWarehouses is an interface that represents a warehouse repository
type RepositoryWarehouses interface {
	// GetOne returns a warehouse by id
	GetOne(id int) (p Warehouse, err error)
	// Store stores a warehouse
	Store(p *Warehouse) (err error)

	GetProductsCount(id *int) (count []WarehouseProductCount, err error)

	// GetAll returns all warehouses
	GetAll() (w []Warehouse, err error)
}
