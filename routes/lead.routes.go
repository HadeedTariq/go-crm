package routes

import "github.com/gin-gonic/gin"

func LeadRoutes(router *gin.RouterGroup) {
	leadApi := router.Group("/leads")
}
