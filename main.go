package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"details": "Trainer-SOT MVP",
		})
	})

	server.Run()
}
