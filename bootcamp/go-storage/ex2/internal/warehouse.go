package internal

// Warehouse is an struct that represents a product
type Warehouse struct {
	// ID is the unique identifier of the product
	ID int
	// Name is the name of the warehouse
	Name      string
	Address   string
	Telephone string
	Capacity  int
}

type WarehouseProductCount struct {
	Name         string
	ProductCount int
}
