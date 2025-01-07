package service

import "app/internal"

// NewCustomersDefault creates new default service for customer entity.
func NewCustomersDefault(rp internal.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

// CustomersDefault is the default service implementation for customer entity.
type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp internal.RepositoryCustomer
}

// GetTop5SpentMost returns the top 5 customers that spent the most
func (s *CustomersDefault) GetTop5SpentMost() (c []internal.CustomerMostSpent, err error) {
	c, err = s.rp.GetTop5SpentMost()
	return
}

// FindAll returns all customers.
func (s *CustomersDefault) FindAll() (c []internal.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

// Save saves the customer.
func (s *CustomersDefault) Save(c *internal.Customer) (err error) {
	err = s.rp.Save(c)
	return
}

// GetTotalGroupedByCondition returns the total of customers grouped by condition
func (s *CustomersDefault) GetTotalGroupedByCondition() (t map[int]float32, err error) {
	t, err = s.rp.GetTotalGroupedByCondition()
	return
}
