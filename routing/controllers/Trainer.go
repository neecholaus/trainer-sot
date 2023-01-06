package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "trainer/sign-up.html", gin.H{
		"title": "Sign Up",
	})
}

func SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "trainer/sign-in.html", gin.H{
		"title": "Sign In",
	})
}

func Home(c *gin.Context) {
	email, _ := c.Get("email")
	exp, _ := c.Get("sessionExpires")
	c.JSON(200, gin.H{
		"message":        "You are logged in as a trainer",
		"email":          email,
		"sessionExpires": exp,
	})
}
