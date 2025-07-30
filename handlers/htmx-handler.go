package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, title string) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"Title": title,
	})
}
