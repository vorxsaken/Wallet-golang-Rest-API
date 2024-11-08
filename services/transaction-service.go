package services

import (
	"fahmi-wallet/database"
	"fahmi-wallet/models"
	"fmt"
)

type TransactionService struct{}

func (s *TransactionService) InitiateTransaction(
	walletID,
	productID,
	transactionTypeID uint,
	amount float64,
	description string,
	productAmount int,
) (*models.Transaction, error) {
	// Retrieve wallet and validate balance
	var wallet models.Wallet
	if err := database.DB.First(&wallet, walletID).Error; err != nil {
		return nil, fmt.Errorf("wallet not found")
	}

	if wallet.Balance < amount {
		return nil, fmt.Errorf("insufficient balance")
	}

	// Retrieve product and validate availability
	var product models.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		return nil, fmt.Errorf("product not found")
	}
	if product.Availability == false || product.Stock == 0 {
		return nil, fmt.Errorf("product not available")
	}

	totalPrice := float64(productAmount) * product.Price

	if totalPrice > amount {
		return nil, fmt.Errorf("insufficient amount: required at least %.2f but received %.2f", totalPrice, amount)
	}

	// Update wallet balance and product stock
	wallet.Balance -= totalPrice
	product.Stock -= productAmount

	if err := database.DB.Save(&wallet).Error; err != nil {
		return nil, fmt.Errorf("failed to update wallet")
	}
	if err := database.DB.Save(&product).Error; err != nil {
		return nil, fmt.Errorf("failed to update product stock")
	}

	// Create transaction record
	transaction := models.Transaction{
		WalletID:          walletID,
		ProductID:         productID,
		Amount:            amount,
		BalanceAfter:      wallet.Balance,
		ProductStock:      product.Stock,
		Description:       description,
		TransactionTypeID: transactionTypeID,
	}

	if err := database.DB.Create(&transaction).Error; err != nil {
		return nil, fmt.Errorf("failed to create transaction")
	}

	return &transaction, nil
}
