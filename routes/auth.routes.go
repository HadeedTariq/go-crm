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
			c.HTML(http.StatusOK, "layout.html", gin.H{
				"title":           "Register User",
				"ContentTemplate": "auth/register.html",
			})
		})
		authApi.POST("/register", auth.RegisterUser)

		// ✅ Login
		authApi.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "layout.html", gin.H{
				"title":           "Login User",
				"ContentTemplate": "auth/login.html",
			})
		})
		authApi.POST("/login", auth.LoginUser)
	}
}
