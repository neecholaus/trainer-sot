package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Clients(c *gin.Context) {
	c.HTML(http.StatusOK, "trainer/clients.html", gin.H{
		"title": "Clients",
	})
}
