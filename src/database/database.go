package database

import (
	"fmt"
	"log"

	"github.com/matheus-gondim/go-cryptocurrencies-election/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection(configs *config.Config) (*gorm.DB, error) {
	strConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		configs.DB_HOST, configs.DB_PORT, configs.DB_USER, configs.DB_DATABASE, configs.DB_PASSWORD,
	)

	db, err := gorm.Open(postgres.Open(strConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
}
