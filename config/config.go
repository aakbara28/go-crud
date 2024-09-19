package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbUser     = "" // fill with your db user
	dbPassword = "" // fill with your pass
	dbName     = "" // fill with your db name
	dbHost     = "localhost"
	dbPort     = "5432"
)

func SetupDB() *gorm.DB {
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}
