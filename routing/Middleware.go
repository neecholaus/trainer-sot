package routing

import (
	"fmt"
	"nicholas/trainer-sot/db"
	"nicholas/trainer-sot/routing/controllers"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateDBConnection() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := db.CreateDbConnection()
		if err != nil {
			c.Abort()
			c.String(500, "Failed to establish DB connection.")
			return
		}

		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Web
		header := c.GetHeader("Authorization")
		// Json (REST)
		cookie, _ := c.Cookie("session")

		var auth string
		if header != "" {
			auth = header
		} else if cookie != "" {
			auth = cookie
		}

		if auth == "" {
			c.Abort()
			c.Redirect(302, "/trainer/sign-in")
			return
		}

		if auth != "" {
			rawToken := strings.Split(auth, " ")[1]

			claims := controllers.TrainerAuthJwtClaims{}

			// todo - replace secret key with env value
			_, err := jwt.ParseWithClaims(rawToken, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("dummy-secret-key"), nil
			})
			if err != nil {
				fmt.Printf("could not parse jwt token: %s", err)
				c.Abort()
				c.JSON(403, gin.H{
					"error": "Could not parse the provided header token",
				})
				return
			}

			c.Set("email", claims.Email)
			c.Set("sessionExpires", claims.ExpiresAt)
		}

		c.Next()
	}
}
