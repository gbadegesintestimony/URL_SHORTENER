package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// This function can be expanded in the future to load other configurations if needed
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
