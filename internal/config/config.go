package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// DB represents the global database connexion
var DB *gorm.DB

// Port is the port used to access the API
var Port string

// DatabaseURL is the database connexion url
var DatabaseURL string

// LoadConfig initializes the environment variables
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env file")
	}

	Port = ":" + os.Getenv("PORT")
	DatabaseURL = os.Getenv("DATABASE_URL")
}
