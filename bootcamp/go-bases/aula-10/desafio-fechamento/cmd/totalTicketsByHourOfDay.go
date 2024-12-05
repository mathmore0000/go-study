package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/spf13/cobra"
)

var totalTicketsByHourCmd = &cobra.Command{
	Use:   "count-by-hour",
	Short: "Obtém o total de tickets por um valor de hora específico",
	Run: func(cmd *cobra.Command, args []string) {
		hour, err := cmd.Flags().GetString("hour")
		if err != nil {
			fmt.Println("Erro ao obter o destino:", err)
			os.Exit(1)
		}
		if hour == "" {
			fmt.Println("Error: Você deve fornecer o destino usando a flag --hour.")
			cmd.Usage()
			os.Exit(1)
		}
		hourInt, err := strconv.Atoi(hour)
		if err != nil {
			fmt.Println("Error: A flag --hour precisa ser um valor inteiro.")
			cmd.Usage()
			os.Exit(1)

		}
		if hourInt < 0 || hourInt > 24 {
			fmt.Println("Error: A flag --hour precisa ter um valor entre 0 e 24.")
			cmd.Usage()
			os.Exit(1)
		}

		totalTicketsByEvening, timeOfDay := tickets.GetTotalTicketsByTime(time.Date(2024, 12, 02, hourInt, 0, 0, 0, time.Local))
		fmt.Printf("Total number of tickets in the %v -> %v\n", timeOfDay, totalTicketsByEvening)
	},
}

func init() {
	rootCmd.AddCommand(totalTicketsByHourCmd)
	totalTicketsByHourCmd.Flags().String("hour", "", "Destino para procurar os tickets")
}
