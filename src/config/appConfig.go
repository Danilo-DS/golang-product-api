package config

import (
	"fmt"
	"log"
	"os"

	dotEnv "github.com/joho/godotenv"
)

// LoadEnvironment load environment variables of the application
func LoadEnvironment() {

	if err := dotEnv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
}

// GetPort return formatted server port
func GetPort() string {
	return fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
}
