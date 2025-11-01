package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func GetOpenRouterAPIKey() string {
	return os.Getenv("OPENROUTER_API_KEY")
}

func GetPort() string {
	return os.Getenv("PORT")
}
