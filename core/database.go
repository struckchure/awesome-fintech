package core

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"awesome.fintech.org/models"
)

func RunMigrations() {
	db, _ := NewDatabaseConnection()

	err := db.AutoMigrate(
		models.Balance{},
		&models.Ledger{},
		&models.Transaction{},
	)
	if err != nil {
		log.Panicf("database migration failed: %s", err)
	}

	log.Println("Migrations Successfully!")
}

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSL_MODE, DB_TIME_ZONE,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("failed to connect database: %s", err)
	}

	return db, nil
}
