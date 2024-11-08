package models

import "github.com/jinzhu/gorm"

// Product represents a product in the system
type Product struct {
	gorm.Model
	Name         string  `json:"name" gorm:"type:varchar(100);not null"`
	Description  string  `json:"description" gorm:"type:varchar(255)"`
	Price        float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	Availability bool    `json:"availability" gorm:"default:true"` // Availability (true = available, false = out of stock)
	Stock        int     `json:"stock"`
}
