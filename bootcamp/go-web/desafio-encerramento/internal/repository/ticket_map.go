package repository

import (
	internal "app/internal/model"
)

// NewRepositoryTicket creates a new repository for tickets in a map
func NewRepositoryTicket(db map[int]internal.TicketAttributes, dbFile string, lastId int) RepositoryTicket {
	return RepositoryTicket{
		dbFile: dbFile,
		db:     db,
		lastId: lastId,
	}
}

// RepositoryTicket implements the repository interface for tickets in a map
type RepositoryTicket struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]internal.TicketAttributes

	dbFile string

	// lastId represents the last id of the ticket
	lastId int
}

// GetAll returns all the tickets
func (r *RepositoryTicket) GetAllCount() int {
	return len(r.db)
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicket) GetTicketsCountByDestinationCountry(country string) (totalTicketsByDest int) {
	totalTicketsByDest = 0
	for _, v := range r.db {
		if v.Country == country {
			totalTicketsByDest++
		}
	}

	return
}
