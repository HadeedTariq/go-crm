package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/controllers/auth"
)

func AuthRoutes(router *gin.RouterGroup) {
	authApi := router.Group("/auth")
	{
		authApi.GET("/")

		authApi.POST("/register", auth.RegisterUser)

		authApi.POST("/login", auth.LoginUser)
	}
}
