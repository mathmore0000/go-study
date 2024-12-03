package tickets_test

import (
	"testing"
	"time"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

var csvFilePath string = "./tickets_test.csv"

func TestGetTotalTickets(t *testing.T) {}

func TestGetTotalTicketsByDestination(t *testing.T) {
	tickets.InitializeTickets(csvFilePath)
	type args struct {
		destination string
	}
	tests := []struct {
		name         string
		args         args
		want         int
		wantErr      bool
		wantedMsgErr string
	}{
		{name: "Testing for Argentina", args: args{destination: "Argentina"}, want: 10, wantErr: false},
		{name: "Testing for Brazil", args: args{destination: "Brazil"}, want: 34, wantErr: false},
		{name: "Testing for China", args: args{destination: "China"}, want: 145, wantErr: false},
		{name: "Testing for Brazil Wrong", args: args{destination: "Brazil Wrong"}, want: 0, wantErr: true, wantedMsgErr: "No tickets on destination Brazil Wrong found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tickets.GetTotalTicketsByDestination(tt.args.destination)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("GetTotalTicketsByDestination() error = %v, wantErr %v", err, tt.wantErr)
					return
				} else if tt.wantedMsgErr != err.Error() {
					t.Errorf("GetTotalTicketsByDestination() error = %v, wantedMsgErr = %v", err, tt.wantedMsgErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("GetTotalTicketsByDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTotalTicketsByTime(t *testing.T) {
	type args struct {
		time time.Time
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{name: "Testing for `noite`", args: args{time: time.Date(2024, 12, 23, 23, 0, 0, 0, time.Local)}, want: 127, want1: "noite"},
		{name: "Testing for `tarde`", args: args{time: time.Date(2024, 12, 20, 14, 0, 0, 0, time.Local)}, want: 225, want1: "tarde"},
		{name: "Testing for `manhã`", args: args{time: time.Date(2024, 12, 0, 8, 0, 0, 0, time.Local)}, want: 203, want1: "manhã"},
		{name: "Testing for `início da manhã`", args: args{time: time.Date(2024, 12, 8, 4, 0, 0, 0, time.Local)}, want: 244, want1: "início da manhã"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tickets.GetTotalTicketsByTime(tt.args.time)
			if got != tt.want {
				t.Errorf("GetTotalTicketsByTime() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetTotalTicketsByTime() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetPercentageByDestination(t *testing.T) {
	type args struct {
		destination string
	}
	tests := []struct {
		name         string
		args         args
		want         float32
		wantErr      bool
		wantedMsgErr string
	}{
		{name: "Testing for Argentina", args: args{destination: "Argentina"}, want: 1.2515645, wantErr: false},
		{name: "Testing for Brazil", args: args{destination: "Brazil"}, want: 4.255319, wantErr: false},
		{name: "Testing for China", args: args{destination: "China"}, want: 18.147684, wantErr: false},
		{name: "Testing for Brazil Wrong", args: args{destination: "Brazil Wrong"}, want: 0, wantErr: true, wantedMsgErr: "No tickets on destination Brazil Wrong found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tickets.GetPercentageByDestination(tt.args.destination)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("GetPercentageByDestination() error = %v, wantErr %v", err, tt.wantErr)
					return
				} else if err.Error() != tt.wantedMsgErr {
					t.Errorf("GetPercentageByDestination() error = %v, wantedMsgErr = %v", err, tt.wantedMsgErr)
				}
			}
			if got != tt.want {
				t.Errorf("GetPercentageByDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
