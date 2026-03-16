package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	OpenAIKey string
	Port      string
	Cors      string
}

func LoadConfig() *Config {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system env")
	}

	key := os.Getenv("OPENAI_API_KEY")
	port := os.Getenv("PORT")
	cors := os.Getenv("CORS")

	if key == "" {
		log.Fatal("OPENAI_API_KEY not set")
	}

	return &Config{
		OpenAIKey: key,
		Port:      port,
		Cors:      cors,
	}
}
