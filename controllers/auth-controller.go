package controllers

import (
	"fahmi-wallet/services"
	"fahmi-wallet/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authService = services.AuthService{}

type AuthController struct{}

// Login function to handle user login
func (s *AuthController) Login(c *gin.Context) {
	var input validations.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Authenticate user using the service
	token, err := authService.AuthenticateUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the JWT token if authentication is successful
	c.JSON(http.StatusOK, gin.H{"token": token})
}
