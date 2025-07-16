package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", controllers.GetUsers)
	}
}
