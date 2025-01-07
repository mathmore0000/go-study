package internal

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	// FindAll returns all customers saved in the database.
	FindAll() (c []Customer, err error)
	// Save saves a customer into the database.
	Save(c *Customer) (err error)
	// GetTotalGroupedByCondition returns the total of customers grouped by condition
	GetTotalGroupedByCondition() (t map[int]float32, err error)
	// GetTop5SpentMost returns the top 5 customers that spent the most
	GetTop5SpentMost() (c []CustomerMostSpent, err error)
}
