package cmd

import (
	"fmt"
	"os"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/spf13/cobra"
)

var csvFilePath string

// rootCmd representa o comando raiz da aplicação
var rootCmd = &cobra.Command{
	Use:   "dfecha",
	Short: "dfecha é uma ferramenta de linha de comando para o desafío final de bases do go",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if csvFilePath == "" {
			fmt.Println("Error: You must provide the CSV file path using the --csv flag.")
			cmd.Usage()
			os.Exit(1)
		}

		fmt.Printf("Bem-vindo ao desafío de fechamento! Usando o arquivo CSV: %s\n", csvFilePath)

		tickets.InitializeTickets(csvFilePath)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Defina o flag csv usado que precisa ser enforçada globalmente
	rootCmd.PersistentFlags().StringVar(&csvFilePath, "csv", "", "Caminho para o arquivo CSV")
}
