package controllers

import (
	"fmt"
	"nicholas/trainer-sot/db"
	"nicholas/trainer-sot/db/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type TrainerAuthJwtClaims struct {
	Email string `json:"email,omitempty"`
	jwt.RegisteredClaims
}

type signUpRestRequestBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func SignUpRest(c *gin.Context) {
	var body signUpRestRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println("error while binding sign up request body")
		c.JSON(500, gin.H{
			"error": "Could not read body.",
		})
		return
	}

	// todo - improved validation (email type, length, etc.)
	if body.Email == "" || body.Password == "" || body.FirstName == "" || body.LastName == "" {
		c.JSON(400, gin.H{
			"error": "Missing required field(s).",
		})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 1)
	if err != nil {
		fmt.Println("error while hashing password during sign up")
		c.JSON(500, gin.H{
			"error": "Could not hash password.",
		})
		return
	}

	trainer := models.Trainer{
		Email:     body.Email,
		Password:  string(hashedPass),
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	// Ensure email is not tied to an existing account
	var existing models.Trainer // Don't actually care about the result, can I avoid passing this somehow?
	res := db.Db.Model(&models.Trainer{}).
		Where("email = ?", trainer.Email).
		Limit(1).
		Find(&existing)
	if res.RowsAffected > 0 {
		fmt.Printf("sign up attempted for email that is already in use: %s\n", trainer.Email)
		c.JSON(400, gin.H{
			"error": "Account already exists with this email",
			"data": gin.H{
				"email": trainer.Email,
			},
		})
		return
	}
	if res.Error != nil {
		fmt.Printf("error while checking for existing trainer account: %s\n", err)
		c.JSON(500, gin.H{
			"error": "Error while checking for existing account.",
		})
		return
	}

	storedTrainer := db.Db.Create(&trainer)
	if storedTrainer.Error != nil {
		fmt.Println("error while storing new trainer sign up")
		c.JSON(500, gin.H{
			"error": "Could not store trainer",
		})
	}

	c.JSON(200, gin.H{
		"message": "Sign up was successful",
	})
}

type signInRestRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// signInRest accepts and validates trainer credentials.
// Then writes a request with a newly generated JWT token set in the Authorization header.
func SignInRest(c *gin.Context) {
	var body signInRestRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println("error while binding sign in request body")
		c.JSON(500, gin.H{
			"error": "Could not complete sign in request.",
		})
		return
	}

	if body.Email == "" {
		fmt.Println("email was not provided in sign up request body")
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
		c.JSON(500, gin.H{
			"error": "Could not generate your auth token.",
		})
		return
	}

	c.Header("Authorization", fmt.Sprintf("Bearer %s", signed))
	c.JSON(200, gin.H{
		"message": "You have been signed in.",
	})
}