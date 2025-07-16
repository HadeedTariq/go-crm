package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.HTML(200, "layout.html", gin.H{
		"Title": "Welcome Page",
		"Name":  "Hadeed",
	})
}
