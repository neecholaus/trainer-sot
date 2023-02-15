package main

import (
	"fmt"
	"nicholas/trainer-sot/db"
	"nicholas/trainer-sot/routing"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error while loading .env file: %s\n", err)
		return
	}

	if os.Getenv("JWT_SECRET_KEY") == "" {
		fmt.Println("no jwt secret key found")
		return
	}

	// Handle one off commands such as db migration
	cmd, isCommand := determineIfCommand()
	if isCommand {
		fmt.Printf("Handling as command: %s\n", cmd)
		return
	}

	server := gin.Default()

	// Register public files
	server.Static("/public", "./resources/public")

	// Db connection ensured on each request
	server.Use(routing.CreateDBConnection())

	routing.RegisterRoutes(server)

	_ = server.Run(":80")
}

func determineIfCommand() (string, bool) {
	if len(os.Args) < 2 {
		return "", false
	}

	cmd := os.Args[1]

	if cmd == "migrate" {
		_, err := db.EnsureDbConnection()
		if err != nil {
			fmt.Println("Db connection failed")
			return "", true
		}
		db.Migrate()
		return cmd, true
	}

	return "", false
}
