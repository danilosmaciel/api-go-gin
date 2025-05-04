package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(c *gin.Context) {
	user := c.ShouldBindBodyWithJSON("user")
	password := c.ShouldBindBodyWithJSON("password")
	fmt.Printf("User: %v\nPassword: %v", user, password)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "OK",
	})
}
