package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
}

func Init() Config {
	config := Config{}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	config.PORT = fmt.Sprintf(":%v", os.Getenv("PORT"))
	if config.PORT == ":" || config.PORT == "" {
		config.PORT = ":5000"
	}

	config.DB_DATABASE = os.Getenv("DB_DATABASE")
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_PORT = os.Getenv("DB_PORT")

	return config
}
