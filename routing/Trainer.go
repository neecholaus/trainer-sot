package routing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type SignInRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TrainerAuthJwtClaims struct {
	Email string `json:"email,omitempty"`
	jwt.RegisteredClaims
}

func SignIn(c *gin.Context) {
	var body SignInRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println("error while binding sign in request body")
		c.String(500, "Could not complete sign in request.")
		return
	}

	if body.Email == "" {
		c.JSON(400, gin.H{
			"error": "Email was not provided",
		})
		return
	}

	// todo - validate credentials before making token

	claims := TrainerAuthJwtClaims{
		body.Email,
		jwt.RegisteredClaims{
			Issuer:    "",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// todo - replace secret key with env value
	signed, err := token.SignedString([]byte("dummy-secret-key"))
	if err != nil {
		fmt.Println("error while signing jwt key")
		c.String(500, "Could not generate your auth token.")
		return
	}

	fmt.Printf("token: %s\n", signed)

	c.Header("Authorization", fmt.Sprintf("Bearer %s", signed))
	c.JSON(200, gin.H{
		"message": "Welcome to the source of truth",
	})
}