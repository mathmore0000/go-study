package service

import "app/internal"

// NewInvoicesDefault creates new default service for invoice entity.
func NewInvoicesDefault(rp internal.RepositoryInvoice) *InvoicesDefault {
	return &InvoicesDefault{rp}
}

// InvoicesDefault is the default service implementation for invoice entity.
type InvoicesDefault struct {
	// rp is the repository for invoice entity.
	rp internal.RepositoryInvoice
}

// UpdateTotal updates the total of a invoice
func (s *InvoicesDefault) UpdateTotal() (err error) {
	err = s.rp.UpdateTotal()
	return
}

// FindAll returns all invoices.
func (s *InvoicesDefault) FindAll() (i []internal.Invoice, err error) {
	i, err = s.rp.FindAll()
	return
}

// Save saves the invoice.
func (s *InvoicesDefault) Save(i *internal.Invoice) (err error) {
	err = s.rp.Save(i)
	return
}
