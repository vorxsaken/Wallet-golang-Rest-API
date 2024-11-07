package services

import (
	"errors"
	"fahmi-wallet/database"
	"fahmi-wallet/models"

	"gorm.io/gorm"
)

// WalletService provides methods for deposit and withdrawal operations
type WalletService struct{}

// CreateWallet creates a new wallet for the user
func (s *WalletService) CreateWallet(userID uint, currency string) (*models.Wallet, error) {
	var existingWallet models.Wallet
	// Check if the user already has a wallet
	if err := database.DB.Where("user_id = ?", userID).First(&existingWallet).Error; err == nil {
		// Return an error if the wallet already exists
		return nil, errors.New("wallet already exists for this user")
	}

	// Create a new wallet
	wallet := models.Wallet{
		UserID:   userID,
		Balance:  0.0,
		Currency: currency,
	}

	// Save the wallet to the database
	if err := database.DB.Create(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

// GetWallet retrieves a wallet by user ID
func (s *WalletService) GetWallet(userID uint) (*models.Wallet, error) {
	var wallet models.Wallet
	// Fetch wallet by user ID
	if err := database.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("wallet not found")
		}
		return nil, err
	}

	return &wallet, nil
}

// UpdateWallet updates the wallet balance or currency
func (s *WalletService) UpdateWallet(userID uint, balance float64, currency string) (*models.Wallet, error) {
	var wallet models.Wallet
	// Fetch wallet by user ID
	if err := database.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("wallet not found")
		}
		return nil, err
	}

	// Update wallet fields
	if balance >= 0 {
		wallet.Balance = balance
	}
	if currency != "" {
		wallet.Currency = currency
	}

	// Save the updated wallet
	if err := database.DB.Save(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

// Deposit function to increase the user's balance
func (s *WalletService) Deposit(userID uint, amount float64, description string) (*models.Wallet, error) {
	var wallet models.Wallet
	var transactionType models.TransactionType

	// Find the wallet by userID
	if err := database.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("wallet not found")
		}
		return nil, err
	}

	// Find the "Deposit" transaction type
	if err := database.DB.Where("type = ?", "Deposit").First(&transactionType).Error; err != nil {
		return nil, errors.New("transaction type 'Deposit' not found")
	}

	// Validate deposit amount
	if amount <= 0 {
		return nil, errors.New("deposit amount must be greater than zero")
	}

	// Update the wallet balance
	wallet.Balance += amount

	// Save the updated wallet
	if err := database.DB.Save(&wallet).Error; err != nil {
		return nil, err
	}

	// Record the transaction
	transaction := models.Transaction{
		WalletID:        wallet.ID,
		Amount:          amount,
		BalanceAfter:    wallet.Balance,
		Description:     description,
		TransactionType: transactionType,
	}

	if err := database.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

// Withdraw function to decrease the user's balance
func (s *WalletService) Withdraw(userID uint, amount float64, description string) (*models.Wallet, error) {
	var wallet models.Wallet
	var transactionType models.TransactionType

	// Find the wallet by userID
	if err := database.DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("wallet not found")
		}
		return nil, err
	}

	// Find the "Withdrawal" transaction type
	if err := database.DB.Where("type = ?", "Withdrawal").First(&transactionType).Error; err != nil {
		return nil, errors.New("transaction type 'Withdrawal' not found")
	}

	// Validate withdrawal amount
	if amount <= 0 {
		return nil, errors.New("withdrawal amount must be greater than zero")
	}

	if wallet.Balance < amount {
		return nil, errors.New("insufficient balance")
	}

	// Deduct the amount from wallet balance
	wallet.Balance -= amount

	// Save the updated wallet
	if err := database.DB.Save(&wallet).Error; err != nil {
		return nil, err
	}

	// Record the transaction
	transaction := models.Transaction{
		WalletID:        wallet.ID,
		Amount:          amount,
		BalanceAfter:    wallet.Balance,
		Description:     description,
		TransactionType: transactionType,
	}

	if err := database.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}
