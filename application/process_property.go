package application

import (
	"github.com/yourusername/intellect/domain"
	"github.com/yourusername/intellect/infrastructure"
	"log"
)

func ProcessProperties(inputCsvPath string) error {
	// Carregar as propriedades do CSV (essa parte já foi implementada)
	properties := []domain.Property{} // Preencha isso com a leitura do CSV e chamadas do Amazon Bedrock

	repo := infrastructure.NewMySQLPropertyRepository()

	for _, property := range properties {
		err := repo.Save(property)
		if err != nil {
			log.Printf("Erro ao salvar propriedade %s: %v", property.ID, err)
		}
	}

	return nil
}
