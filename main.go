package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/config"
	"github.com/hadeedtariq/go-crm/handlers"
	"github.com/hadeedtariq/go-crm/routes"
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

	validators.InitValidator()

	api := router.Group("/api")

	api.GET("/", handlers.Home)
	api.GET("/partials", handlers.Partial)

	routes.AuthRoutes(api)
	routes.LeadRoutes(api)

	router.Run(":3000")
}
