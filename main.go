package main

import (
	"github.com/gin-gonic/gin"
	"nicholas/trainer-sot/routing"
)

func main() {
	server := gin.Default()

	// Db connection created for each request
	server.Use(routing.CreateDBConnection())

	routing.RegisterRoutes(server)

	_ = server.Run(":80")
}
