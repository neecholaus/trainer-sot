package routing

import (
	"fmt"
	"nicholas/trainer-sot/db"
	"nicholas/trainer-sot/routing/controllers"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateDBConnection() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := db.EnsureDbConnection()
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
		// Json (backup when no session cookie available for fetch)
		header := c.GetHeader("Authorization")
		// Web (primary)
		cookie, _ := c.Cookie("session")

		var auth string
		if header != "" {
			auth = header
		} else if cookie != "" {
			auth = cookie
		}

		if auth == "" {
			handleNoAuth(c)
			return
		}

		if auth != "" {
			rawToken := strings.Split(auth, " ")[1]

			claims := controllers.TrainerAuthJwtClaims{}

			_, err := jwt.ParseWithClaims(rawToken, &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			})
			if err != nil {
				fmt.Printf("could not parse jwt token: %s", err)
				handleNoAuth(c)
				return
			}

			c.Set("trainerId", claims.TrainerId)
			c.Set("email", claims.Email)
		}

		c.Next()
	}
}

// handleNoAuth
// Will handle response by returning response that is appropriate to the request. (json or redirect)
func handleNoAuth(c *gin.Context) {
	c.Abort()
	if c.GetHeader("Content-Type") == "application/json" {
		c.JSON(403, gin.H{
			"error": "You are not authenticated.",
		})
		return
	}

	c.Redirect(302, "/trainer/sign-in")
}
