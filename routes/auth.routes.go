package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/controllers/auth"
)

func AuthRoutes(router *gin.Engine) {
	authApi := router.Group("/api/auth")
	{
		authApi.GET("/")
		authApi.POST("/register", auth.RegisterUser)
		authApi.POST("/login")
	}
}
