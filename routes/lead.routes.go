package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/controllers/leads"
)

func LeadRoutes(router *gin.RouterGroup) {
	// leadsApi := router.Group("/leads")
	// {

	// }

	contacts := router.Group("/contacts")
	{
		contacts.POST("/create", leads.CreateContact)
		contacts.GET("/create", func(c *gin.Context) {
			c.HTML(http.StatusOK, "layout.html", gin.H{
				"Title":           "Create New Contact",
				"ContentTemplate": "leads/create-contact.html",
			})
		})
	}

}
