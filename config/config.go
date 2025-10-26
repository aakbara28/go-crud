package config

import (
	"fmt"
	"go-crud/internal/environment"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	config := environment.AppConfig.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.Name, config.Port, config.SSLMode)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	// Always close database connection when done
	dbSql, _ := db.DB()
	dbSql.SetMaxOpenConns(25)
	dbSql.SetMaxIdleConns(25)
	
	return db
}
