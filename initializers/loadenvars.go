package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}
