package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetOpenRouterAPIKey() string {
	return os.Getenv("OPENROUTER_API_KEY")
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetAllowedOrigin() string {
	return os.Getenv("ALLOWED_ORIGIN")
}
