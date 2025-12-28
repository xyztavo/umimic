package config

import (
	"fmt"
	"os"
	"strings"

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

func GetAllowedOrigins() []string {
	origins := os.Getenv("ALLOWED_ORIGINS")
	// Remove aspas (caso estejam no .env)
	origins = strings.Trim(origins, `"'`)

	parts := strings.Split(origins, ",")

	return parts
}

func GetRedisURL() string {
	return os.Getenv("UPSTASH_REDIS_REST_URL")
}

func GetRedisToken() string {
	return os.Getenv("UPSTASH_REDIS_REST_TOKEN")
}
