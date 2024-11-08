package models

import "github.com/jinzhu/gorm"

// TransactionType represents types of transactions (Deposit, Withdrawal, etc.)
type TransactionType struct {
	gorm.Model
	Type        string `json:"type"`
	Description string `json:"description"`
}

// Transaction represents a wallet transaction (Deposit, Withdrawal, etc.)
type Transaction struct {
	gorm.Model
	WalletID          uint            `json:"wallet_id"`
	ProductID         uint            `json:"product_id"`
	Amount            float64         `json:"amount"`
	BalanceAfter      float64         `json:"balance_after"`
	Description       string          `json:"description"`
	ProductStock      int             `json:"product_stock"`
	TransactionTypeID uint            `json:"transaction_type_id"`
	TransactionType   TransactionType `gorm:"foreignkey:TransactionTypeID"`
}
