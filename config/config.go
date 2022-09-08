package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	Database   Database
}

type Database struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var config Config
	config.ServerPort = os.Getenv("SERVER_PORT")
	config.Database.DBHost = os.Getenv("DB_HOST")
	config.Database.DBPort = os.Getenv("DB_PORT")
	config.Database.DBUser = os.Getenv("DB_USERNAME")
	config.Database.DBPassword = os.Getenv("DB_PASSWORD")
	config.Database.DBName = os.Getenv("DB_DATABASE")
	return config
}
