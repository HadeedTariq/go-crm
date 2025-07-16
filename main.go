package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/config"
	"github.com/hadeedtariq/go-crm/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectToDatabase()

	router := gin.Default()
	templates := template.Must(template.ParseGlob("templates/**/*"))

	router.SetHTMLTemplate(templates)

	routes.RegisterRoutes(router)

	router.Run(":3000")
}
