package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
