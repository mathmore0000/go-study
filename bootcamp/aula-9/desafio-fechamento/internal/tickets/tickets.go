package tickets

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bootcamp-go/desafio-go-bases/internal/csv"
)

type Ticket struct {
	Name    string
	Email   string
	Country string
	Time    time.Time
	Cost    float32
}

var Tickets []Ticket

var TicketsByCountry map[string][]Ticket = map[string][]Ticket{}
var TicketsByTimeOfDay map[string][]Ticket = map[string][]Ticket{}

func createTicket(name string, email string, country string, time time.Time, cost float32) (ticket *Ticket) {
	return &Ticket{name, email, country, time, cost}
}

func InitializeTickets(ticketsCsvFilePath string) {
	records := csv.ReadFile(ticketsCsvFilePath)
	for _, record := range records {
		var splittedTime []string = strings.Split(record[4], ":")

		newTicketHour, err := strconv.Atoi(splittedTime[0])
		treatErrorConversion(err)

		newTicketMinute, err := strconv.Atoi(splittedTime[1])
		treatErrorConversion(err)

		newTicketCost, err := strconv.ParseFloat(record[5], 32)
		treatErrorConversion(err)

		newTicket := createTicket(record[1], record[2], record[3], time.Date(2024, 12, 02, newTicketHour, newTicketMinute, 0, 0, time.Local), float32(newTicketCost))
		newTicketTimeOfDay := getTimeOfDayByTime(newTicket.Time)

		Tickets = append(Tickets, *newTicket)
		TicketsByCountry[newTicket.Country] = append(TicketsByCountry[newTicket.Country], *newTicket)
		TicketsByTimeOfDay[newTicketTimeOfDay] = append(TicketsByTimeOfDay[newTicketTimeOfDay], *newTicket)
	}
}

func GetTotalTicketsByDestination(destination string) (int, error) {
	if _, ok := TicketsByCountry[destination]; !ok {
		return 0, errors.New(fmt.Sprintf("No tickets on destination %v found", destination))
	}
	return len(TicketsByCountry[destination]), nil
}

func GetTotalTicketsByTime(time time.Time) (int, string) {
	var timeOfDay string = getTimeOfDayByTime(time)
	return len(TicketsByTimeOfDay[timeOfDay]), timeOfDay
}

func GetPercentageByDestination(destination string) (float32, error) {
	if _, ok := TicketsByCountry[destination]; !ok {
		return 0, errors.New(fmt.Sprintf("No tickets on destination %v found", destination))
	}
	return (float32(len(TicketsByCountry[destination])) / float32(len(Tickets))) * 100, nil
}

func treatErrorConversion(err error) {
	if err != nil {
		panic(err)
	}
}

func getTimeOfDayByTime(time time.Time) (timeOfDay string) {
	var hour int = time.Hour()
	if hour <= 6 {
		return "início da manhã"
	}
	if hour <= 12 {
		return "manhã"
	}
	if hour <= 19 {
		return "tarde"
	}
	if hour <= 23 {
		return "noite"
	}
	return "desconhecido"
}
