package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/controllers/leads"
)

func LeadRoutes(router *gin.RouterGroup) {
	leadApi := router.Group("/leads")
	{
		leadApi.POST("/contact/create", leads.CreateContact)
	}

}
