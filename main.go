package main

import (
	"github.com/gin-gonic/gin"
	"nicholas/trainer-sot/routing"
)

func main() {
	server := gin.Default()

	routing.RegisterRoutes(server)

	_ = server.Run(":80")
}
