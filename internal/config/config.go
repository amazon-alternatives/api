package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Port string
var DatabaseUrl string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env file")
	}

	Port = ":" + os.Getenv("PORT")
	DatabaseUrl = os.Getenv("DATABASE_URL")
}
