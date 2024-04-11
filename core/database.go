package core

import (
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
		log.Panicln(err)
	}

	log.Println("Migrations Successfully!")
}

func NewDatabaseConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Africa/Lagos"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicln("failed to connect database")
	}

	return db, nil
}
