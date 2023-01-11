package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nicholas/trainer-sot/db"
	"nicholas/trainer-sot/db/models"
)

// ClientListRest
// Endpoint enabling front end to "search" for clients. By default, returns the last `x`
// amount sorting by newest first.
func ClientListRest(c *gin.Context) {
	trainerId, _ := c.Get("trainerId")
	if trainerId == nil {
		c.JSON(400, gin.H{
			"error": "No trainer info found.",
		})
		return
	}

	var clients []models.Client
	_ = db.Db.Model(models.Client{}).
		Where("trainer_id = ?", trainerId).
		Limit(10).
		Find(&clients)

	c.JSON(http.StatusOK, gin.H{
		"trainerId": trainerId,
		"clients":   clients,
	})
}
