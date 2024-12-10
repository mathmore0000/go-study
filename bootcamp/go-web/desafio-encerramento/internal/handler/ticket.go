package handler

import (
	"app/internal/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandlerTicketDefault(sv *service.ServiceTicketDefault) *HandlerTicketDefault {
	return &HandlerTicketDefault{
		sv: sv,
	}
}

type HandlerTicketDefault struct {
	sv *service.ServiceTicketDefault
}

func (h *HandlerTicketDefault) GetPercentageTicketsByDestinationCountry(w http.ResponseWriter, r *http.Request) {
	// get destination country
	dest := chi.URLParam(r, "dest")
	// get average
	percentage, err := h.sv.GetPercentageTicketsByDestinationCountry(dest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write response
	w.Write([]byte(fmt.Sprintf("%.2f", percentage)))
}

func (h *HandlerTicketDefault) GetTicketsCountByDestinationCountry(w http.ResponseWriter, r *http.Request) {
	dest := chi.URLParam(r, "dest")
	w.Write([]byte(fmt.Sprintf("%d", h.sv.GetTicketsCountByDestinationCountry(dest))))
}
