package service

import (
	"errors"

	"app/internal/model"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp model.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp model.RepositoryTicket) ServiceTicketDefault {
	return ServiceTicketDefault{
		rp: rp,
	}
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(destination string) (average float32, err error) {
	totalTicketsByDest := s.rp.GetTicketsCountByDestinationCountry(destination)
	if totalTicketsByDest == 0 {
		return 0, errors.New("destination not found")
	}
	return (float32(totalTicketsByDest) / float32(s.rp.GetAllCount())) * 100, nil
}

func (s *ServiceTicketDefault) GetTicketsCountByDestinationCountry(destination string) int {
	return s.rp.GetTicketsCountByDestinationCountry(destination)
}
