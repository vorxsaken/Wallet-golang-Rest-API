package services

import (
	"errors"
	"fahmi-wallet/auth"
	"fahmi-wallet/database"
	"fahmi-wallet/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (s *AuthService) AuthenticateUser(email, password string) (string, error) {
	var user models.User

	// Find user by email
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	// Check if the provided password matches the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token upon successful authentication
	token, err := auth.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
