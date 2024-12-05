package cmd

import (
	"fmt"
	"os"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/spf13/cobra"
)

var totalTicketsByDestinationCmd = &cobra.Command{
	Use:   "count-by-destination",
	Short: "Obtém o total de tickets para um destino específico",
	Run: func(cmd *cobra.Command, args []string) {
		destination, err := cmd.Flags().GetString("destination")
		if err != nil {
			fmt.Println("Erro ao obter o destino:", err)
			os.Exit(1)
		}

		if destination == "" {
			fmt.Println("Error: Você deve fornecer o destino usando a flag --destination.")
			cmd.Usage()
			os.Exit(1)
		}

		total, err := tickets.GetTotalTicketsByDestination(destination)
		if err != nil {
			fmt.Println("Erro ao obter o total de tickets:", err)
			os.Exit(1)
		}

		fmt.Printf("Total de tickets para %s: %d\n", destination, total)
	},
}

func init() {
	rootCmd.AddCommand(totalTicketsByDestinationCmd)
	totalTicketsByDestinationCmd.Flags().String("destination", "", "Destino para procurar os tickets")
}
