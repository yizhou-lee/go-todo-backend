package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMySQLURI() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MYSQLURI")
}
