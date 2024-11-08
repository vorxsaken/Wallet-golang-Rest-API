package services

import (
	"errors"
	"fahmi-wallet/database"
	"fahmi-wallet/models"

	"gorm.io/gorm"
)

type ProductService struct{}

// Create a new product
func (s *ProductService) CreateProduct(name, description string, price float64, availability bool, stock int) (*models.Product, error) {
	product := models.Product{
		Name:         name,
		Description:  description,
		Price:        price,
		Availability: availability,
		Stock:        stock,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

// Get all products
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// Get a product by ID
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

// Update a product
func (s *ProductService) UpdateProduct(id uint, name, description string, price float64, availability bool, stock int) (*models.Product, error) {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return nil, err
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.Availability = availability
	product.Stock = stock

	if err := database.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

// Delete a product
func (s *ProductService) DeleteProduct(id uint) error {
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
