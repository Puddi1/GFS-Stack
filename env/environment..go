package env

import (
	"log"

	"github.com/joho/godotenv"
)

var ENVs map[string]string

func Init_env() {
	var err error
	ENVs, err = godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
