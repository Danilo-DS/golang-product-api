package config

import (
	"fmt"
	"log"
	"os"

	dotEnv "github.com/joho/godotenv"
)

// get .env https://stackoverflow.com/questions/66518873/how-to-read-env-files-by-runnnig-go-application
func LoadEnvironment() {

	if err := dotEnv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
}

func GetPort() string {
	return fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
}
