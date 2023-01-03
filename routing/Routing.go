package routing

import "github.com/gin-gonic/gin"

func RegisterRoutes(c *gin.Engine) {
	// Trainer
	trainerNonAuthGroup := c.Group("/trainer")
	{
		trainerNonAuthGroup.GET("/sign-in", signIn)
		trainerNonAuthGroup.POST("/sign-in", signInRest)
	}
	// Trainer Authed
	trainerAuthGroup := c.Group("/trainer", Auth())
	{
		trainerAuthGroup.GET("/", home)
	}
}
