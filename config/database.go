package config

import (
	"log"

	"github.com/hadeedtariq/go-crm/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := utils.GetEnv("DB_CONN_STRING", "kaka")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// err = database.AutoMigrate(
	// 	&models.User{},
	// 	&models.Deal{},
	// 	&models.Account{},
	// 	&models.Activity{},
	// 	&models.Contact{},
	// )

	// if err != nil {
	// 	log.Fatal("Failed to auto-migrate:", err)
	// }
	DB = database
}
