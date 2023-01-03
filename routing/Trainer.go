package routing

import "github.com/gin-gonic/gin"

func signIn(c *gin.Context) {
	c.String(200, "This will be your login page.")
}

func home(c *gin.Context) {
	email, _ := c.Get("email")
	exp, _ := c.Get("sessionExpires")
	c.JSON(200, gin.H{
		"message":        "You are logged in as a trainer",
		"email":          email,
		"sessionExpires": exp,
	})
}
