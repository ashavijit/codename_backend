package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMONGO_URI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("MONGO_URI")
}

func Redis_DB() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("REDIS_URL")
} 