package repository

// NewRepositoryTicketMock creates a new repository for tickets in a map
func NewRepositoryTicketMock() *RepositoryTicketMock {
	return &RepositoryTicketMock{}
}

// RepositoryTicketMock implements the repository interface for tickets
type RepositoryTicketMock struct {
	// FuncGet represents the mock for the Get function
	FuncGetAllCount func() int
	// FuncGetTicketsByDestinationCountry
	FuncGetTicketsCountByDestinationCountry func(country string) int

	// Spy verifies if the methods were called
	Spy struct {
		// Get represents the spy for the Get function
		Get int
		// GetTicketsByDestinationCountry represents the spy for the GetTicketsByDestinationCountry function
		GetTicketsCountByDestinationCountry int
	}
}

// GetAll returns all the tickets
func (r *RepositoryTicketMock) GetAllCount() int {
	// spy
	r.Spy.Get++

	// mock
	return r.FuncGetAllCount()
}

func (r *RepositoryTicketMock) GetTicketsCountByDestinationCountry(country string) int {
	// spy
	r.Spy.GetTicketsCountByDestinationCountry++

	// mock
	return r.FuncGetTicketsCountByDestinationCountry(country)
}
