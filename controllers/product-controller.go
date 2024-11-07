package controllers

import (
	"fahmi-wallet/services"
	"fahmi-wallet/validations"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

var productService = services.ProductService{}

// Create Product
func (s *ProductController) CreateProduct(c *gin.Context) {
	var input validations.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Create product using the service
	product, err := productService.CreateProduct(input.Name, input.Description, input.Price, input.Availability)
	if err != nil {
		log.Println("Error creating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// Get All Products
func (s *ProductController) GetAllProducts(c *gin.Context) {
	products, err := productService.GetAllProducts()
	if err != nil {
		log.Println("Error fetching products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Get Product by ID
func (s *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	productIDUint, _ := strconv.ParseUint(id, 10, 32)
	product, err := productService.GetProductByID(uint(productIDUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Update Product
func (s *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input validations.UpdateProductInput

	// Bind JSON body to struct and validate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	productIDUint, _ := strconv.ParseUint(id, 10, 32)

	// Update product using the service
	product, err := productService.UpdateProduct(uint(productIDUint), input.Name, input.Description, input.Price, input.Availability)
	if err != nil {
		log.Println("Error updating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Delete Product
func (s *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productIDUint, _ := strconv.ParseUint(id, 10, 32)
	if err := productService.DeleteProduct(uint(productIDUint)); err != nil {
		log.Println("Error deleting product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Product deleted successfully"})
}
