package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

var strConn string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_USER := os.Getenv("DB_USER")
	DB_PORT := os.Getenv("DB_PORT")

	strConn = "postgres://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_DATABASE + "?sslmode=disable"
}

func main() {
	m, err := migrate.New("file://migration/sql", strConn)
	if err != nil {
		fmt.Println("=> ", err)
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
