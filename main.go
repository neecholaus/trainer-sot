package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"details": "Trainer-SOT MVP",
		})
	})

	trainerNS := server.Group("trainer", auth())

	trainerNS.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You are logged in as a trainer",
		})
	})

	_ = server.Run(":80")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "" {
			// todo - validate provided token
		} else {
			c.Abort()
			c.JSON(403, gin.H{
				"error": "You are not logged in",
			})
		}
		c.Next()
	}
}
