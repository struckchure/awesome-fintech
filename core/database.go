package core

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"awesome.fintech.org/core/constants"
	"awesome.fintech.org/models"
)

func RunMigrations() {
	db, _ := NewDatabaseConnection(constants.NewEnv())

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

func NewDatabaseConnection(env *constants.Env) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		env.DB_HOST, env.DB_USER, env.DB_PASSWORD, env.DB_NAME, env.DB_PORT, env.DB_SSL_MODE, env.DB_TIME_ZONE,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("failed to connect database: %s", err)
	}

	return db, nil
}
