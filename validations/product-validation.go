// validations/product.go
package validations

// CreateProductInput struct to receive product creation requests
type CreateProductInput struct {
	Name         string  `json:"name" binding:"required,min=3,max=100"`
	Description  string  `json:"description" binding:"max=255"`
	Price        float64 `json:"price" binding:"required,gt=0"`
	Availability bool    `json:"availability" binding:"required"`
	Stock        int     `json:"stock" binding:"required"`
}

// UpdateProductInput struct to receive product update requests
type UpdateProductInput struct {
	Name         string  `json:"name" binding:"required,min=3,max=100"`
	Description  string  `json:"description" binding:"max=255"`
	Price        float64 `json:"price" binding:"required,gt=0"`
	Availability bool    `json:"availability"`
	Stock        int     `json:"stock" binding:"required"`
}
