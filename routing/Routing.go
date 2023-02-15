package routing

import (
	"nicholas/trainer-sot/routing/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Trainer
	trainerNonAuthGroup := server.Group("/trainer")
	{
		trainerNonAuthGroup.GET("/sign-up", controllers.SignUp)
		trainerNonAuthGroup.POST("/sign-up", controllers.SignUpRest)
		trainerNonAuthGroup.GET("/sign-in", controllers.SignIn)
		trainerNonAuthGroup.POST("/sign-in", controllers.SignInRest)
	}
	// Trainer Authed
	trainerAuthGroup := server.Group("/trainer", Auth())
	{
		trainerAuthGroup.GET("", controllers.Home)
		trainerAuthGroup.GET("/clients", controllers.Clients)
		trainerAuthGroup.GET("/clients/list", controllers.ClientListRest)
		trainerAuthGroup.POST("/clients", controllers.ClientCreateRest)
	}
}

func RegisterTemplates(server *gin.Engine) {
	server.LoadHTMLGlob("./resources/views/**/*.html")
}
