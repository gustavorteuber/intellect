package config

import (
	"log"
	"github.com/joho/godotenv"
	"os"
)

// Configura as variáveis de ambiente a partir do arquivo .env
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

// Funções para pegar as variáveis de ambiente
func GetDBUser() string {
	return os.Getenv("DB_USER")
}

func GetDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func GetDBHost() string {
	return os.Getenv("DB_HOST")
}

func GetDBPort() string {
	return os.Getenv("DB_PORT")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func GetAWSRegion() string {
	return os.Getenv("AWS_REGION")
}

func GetAWSAccessKeyID() string {
	return os.Getenv("AWS_ACCESS_KEY_ID")
}

func GetAWSSecretAccessKey() string {
	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}
