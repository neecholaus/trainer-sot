package routing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"nicholas/trainer-sot/db"
	"strings"
)

func CreateDBConnection() gin.HandlerFunc {
	return func(c *gin.Context) {
		exists := db.CreateDbConnection()
		if !exists {
			c.Abort()
			c.String(500, "Failed to establish DB connection.")
			return
		}

		c.Next()
	}
}

func Auth() gin.HandlerFunc {
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

			claims := TrainerAuthJwtClaims{}

			// todo - replace secret key with env value
			_, err := jwt.ParseWithClaims(rawToken, &claims, func(token *jwt.Token) (interface{}, error) {
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

			c.Set("email", claims.Email)
			c.Set("sessionExpires", claims.ExpiresAt)
		}

		c.Next()
	}
}
