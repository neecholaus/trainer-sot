package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"nicholas/trainer-sot/routing"
	"strings"
)

func main() {
	server := gin.Default()

	server.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"details": "Trainer-SOT MVP",
		})
	})

	server.POST("/trainer/sign-in", routing.SignIn)

	trainerNS := server.Group("trainer", auth())

	trainerNS.GET("/", func(c *gin.Context) {
		email, _ := c.Get("email")
		c.JSON(200, gin.H{
			"message": "You are logged in as a trainer",
			"email":   email,
		})
	})

	_ = server.Run(":80")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.Abort()
			c.JSON(403, gin.H{
				"error": "You are not logged in",
			})
			return
		}

		if auth != "" {
			rawToken := strings.Split(auth, " ")[1]

			claims := routing.TrainerAuthJwtClaims{}

			// todo - replace secret key with env value
			token, err := jwt.ParseWithClaims(rawToken, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("dummy-secret-key"), nil
			})
			if err != nil {
				fmt.Printf("could not parse jwt token: %s", err)
				c.Abort()
				c.JSON(403, gin.H{
					"error": "Could not parse the provided auth token",
				})
				return
			}

			fmt.Printf("token: %s\n", token.Claims)
			c.Set("email", claims.Email)
		}

		c.Next()
	}
}
