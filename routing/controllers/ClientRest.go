package controllers

import (
	"fmt"
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

type clientCreateRestRequestBody struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
}

func ClientCreateRest(c *gin.Context) {
	trainerId := c.GetUint("trainerId")
	if trainerId == 0 {
		c.JSON(400, gin.H{
			"error": "No trainer info was found.",
		})
		return
	}

	var body clientCreateRestRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Could not read body.",
		})
		return
	}

	client := models.Client{
		TrainerId: trainerId,
		Email:     body.Email,
		Phone:     body.Phone,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}
	storedClient := db.Db.Create(&client)
	if storedClient.Error != nil {
		fmt.Println("error while storing new client")
		c.JSON(500, gin.H{
			"error":  "Could not store the client.",
			"detail": storedClient.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Client has been created.",
	})
}
