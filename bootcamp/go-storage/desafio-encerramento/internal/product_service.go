package internal

// ServiceProduct is the interface that wraps the basic Product methods.
type ServiceProduct interface {
	// FindAll returns all products.
	FindAll() (p []Product, err error)
	// Save saves a product.
	Save(p *Product) (err error)

	// GetTop5MostSold returns the top 5 most sold products
	GetTop5MostSold() (p []ProductQuantity, err error)
}
