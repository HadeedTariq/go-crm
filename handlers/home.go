package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hadeedtariq/go-crm/views"
)

func Home(c *gin.Context) {
	views.Home().Render(c.Writer)
}

func Partial(c *gin.Context) {
	views.Partial().Render(c.Writer)
}
