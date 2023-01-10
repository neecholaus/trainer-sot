package routing

import (
	"nicholas/trainer-sot/routing/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(c *gin.Engine) {
	// Trainer
	trainerNonAuthGroup := c.Group("/trainer")
	{
		trainerNonAuthGroup.GET("/sign-up", controllers.SignUp)
		trainerNonAuthGroup.POST("/sign-up", controllers.SignUpRest)
		trainerNonAuthGroup.GET("/sign-in", controllers.SignIn)
		trainerNonAuthGroup.POST("/sign-in", controllers.SignInRest)
	}
	// Trainer Authed
	trainerAuthGroup := c.Group("/trainer", Auth())
	{
		trainerAuthGroup.GET("", controllers.Home)
		trainerAuthGroup.GET("/clients", controllers.Clients)
	}
}
