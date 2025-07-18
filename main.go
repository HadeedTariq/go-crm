package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/config"
	"github.com/hadeedtariq/go-crm/routes"
	"github.com/hadeedtariq/go-crm/utils"
	"github.com/hadeedtariq/go-crm/validators"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDatabase()

	router := gin.Default()

	router.SetHTMLTemplate(utils.LoadTemplates("templates"))
	validators.InitValidator()

	routes.AuthRoutes(router)

	router.Run(":3000")
}
