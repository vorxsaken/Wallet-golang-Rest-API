package models

import "github.com/jinzhu/gorm"

// Wallet represents the user's wallet in the system.
type Wallet struct {
	gorm.Model
	UserID       uint    `json:"user_id"`
	Balance      float64 `json:"balance"`
	Currency     string  `json:"currency"`
	User         User    `gorm:"foreignkey:UserID"`
	Transactions []Transaction
}
