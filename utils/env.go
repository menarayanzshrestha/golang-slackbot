package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GetEnvWithKey : get env value of provided key
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

//LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}

}
