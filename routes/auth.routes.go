package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/controllers/auth"
)

func AuthRoutes(router *gin.RouterGroup) {
	authApi := router.Group("/auth")
	{
		authApi.GET("/")

		authApi.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "auth/register.html", gin.H{
				"title": "Register User",
			})
		})
		authApi.POST("/register", auth.RegisterUser)

		// ✅ Login
		authApi.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "auth/login.html", gin.H{
				"title": "Login User",
			})
		})
		authApi.POST("/login", auth.LoginUser)
	}
}
