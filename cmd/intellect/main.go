package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/gustavorteuber/intellect/application"
	"github.com/gustavorteuber/intellect/config"
)

// Variáveis globais
var inputFile string

func main() {
	config.LoadEnvVariables()

	var rootCmd = &cobra.Command{
		Use:   "intellect",
		Short: "Intellect é uma ferramenta de processamento de leilões",
		Run: func(cmd *cobra.Command, args []string) {
			if inputFile == "" {
				fmt.Println("Por favor, forneça o caminho para o arquivo CSV com o flag '--file'.")
				os.Exit(1)
			}

			// Processa o arquivo CSV
			err := application.ProcessProperties(inputFile)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Processamento concluído com sucesso!")
		},
	}

	// Adicionar a flag para selecionar o arquivo
	rootCmd.PersistentFlags().StringVarP(&inputFile, "file", "f", "", "Caminho para o arquivo CSV")

	// Executar o comando CLI
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
