package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/controllers/leads"
	"github.com/hadeedtariq/go-crm/handlers"
)

func LeadRoutes(router *gin.RouterGroup) {
	// leadsApi := router.Group("/leads")
	// {

	// }

	contacts := router.Group("/contacts")
	{
		contacts.POST("/create", leads.CreateContact)
		contacts.GET("/create", func(c *gin.Context) {
			handlers.Render(c, "Create Contact")
		})
	}

}
