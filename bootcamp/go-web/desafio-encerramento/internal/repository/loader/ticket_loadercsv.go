package loader

import (
	internal "app/internal/model"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (map[int]internal.TicketAttributes, error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return nil, err
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	ticketAttributesMap := make(map[int]internal.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", err)
			return nil, err
		}

		// serialize the record
		id, err := strconv.Atoi(record[0])
		if err != nil {
			panic("unable to load tickets, invalid id")
		}

		price, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			panic("unable to load tickets, invalid price")
		}

		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}

		// add the ticket to the map
		ticketAttributesMap[id] = ticket
	}

	return ticketAttributesMap, nil
}
