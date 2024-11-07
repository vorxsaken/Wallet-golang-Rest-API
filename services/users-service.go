package services

import (
	"errors"

	"fahmi-wallet/database"
	"fahmi-wallet/models"
)

type UsersService struct{}

// GetAllUsers retrieves all users from the database.
func (s *UsersService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID retrieves a user by ID from the database.
func (s *UsersService) GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

// CreateUser creates a new user in the database.
func (s *UsersService) CreateUser(username, email, password, role string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Email:    email,
		Password: password, // Plain password; will be hashed by BeforeSave
		Role:     role,
	}

	// Attempt to save the user to the database
	if err := database.DB.Create(user).Error; err != nil {
		return nil, errors.New("could not create user")
	}

	return user, nil
}

// UpdateUser updates an existing user's details.
func (s *UsersService) UpdateUser(user *models.User) error {
	// Check if the user exists by ID before updating
	existingUser := &models.User{}
	if err := database.DB.First(existingUser, user.ID).Error; err != nil {
		return err // Return error if user not found
	}

	// Update the user fields
	if err := database.DB.Model(existingUser).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user by ID from the database.
func (s *UsersService) DeleteUser(id uint) error {
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
