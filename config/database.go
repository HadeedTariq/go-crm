package config

import (
	"log"
	"os"

	"github.com/hadeedtariq/go-crm/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := getEnv("DB_CONN_STRING", "kaka")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = database.AutoMigrate(
		&models.User{},
		&models.Deal{},
		&models.Account{},
		&models.Activity{},
		&models.Contact{},
	)

	if err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}
	DB = database
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
