package controllers

import (
	"fahmi-wallet/services"
	"fahmi-wallet/validations"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var walletService = services.WalletService{}

type WalletController struct{}

// CreateWallet endpoint to create a wallet for the user
func (s *WalletController) CreateWallet(c *gin.Context) {
	userID := c.Param("user_id")
	var input validations.CreateWalletInput

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert user_id to uint
	userIDUint, _ := strconv.ParseUint(userID, 10, 32)

	// Call service to create wallet
	wallet, err := walletService.CreateWallet(uint(userIDUint), input.Currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with created wallet
	c.JSON(http.StatusOK, gin.H{
		"message": "Wallet created successfully",
		"wallet":  wallet,
	})
}

// GetWallet endpoint to retrieve the wallet by user ID
func (s *WalletController) GetWallet(c *gin.Context) {
	userID := c.Param("user_id")

	// Convert user_id to uint
	userIDUint, _ := strconv.ParseUint(userID, 10, 32)

	// Call service to get wallet
	wallet, err := walletService.GetWallet(uint(userIDUint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with wallet details
	c.JSON(http.StatusOK, gin.H{
		"wallet": wallet,
	})
}

// UpdateWallet endpoint to update wallet details
func (s *WalletController) UpdateWallet(c *gin.Context) {
	userID := c.Param("user_id")
	var input validations.UpdateWalletInput

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert user_id to uint
	userIDUint, _ := strconv.ParseUint(userID, 10, 32)

	// Call service to update wallet
	wallet, err := walletService.UpdateWallet(uint(userIDUint), input.Balance, input.Currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with updated wallet details
	c.JSON(http.StatusOK, gin.H{
		"message": "Wallet updated successfully",
		"wallet":  wallet,
	})
}

// Deposit endpoint to add funds to a user's wallet
func (s *WalletController) Deposit(c *gin.Context) {
	userID := c.Param("user_id")
	var input validations.DepositInput

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert user_id to uint
	userIDUint, _ := strconv.ParseUint(userID, 10, 32)

	// Call deposit service
	wallet, err := walletService.Deposit(uint(userIDUint), input.Amount, input.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with updated wallet balance
	c.JSON(http.StatusOK, gin.H{
		"message": "Deposit successful",
		"balance": wallet.Balance,
	})
}

// Withdraw endpoint to deduct funds from a user's wallet
func (s *WalletController) Withdraw(c *gin.Context) {
	userID := c.Param("user_id")
	var input validations.WithdrawInput

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert user_id to uint
	userIDUint, _ := strconv.ParseUint(userID, 10, 32)

	// Call withdraw service
	wallet, err := walletService.Withdraw(uint(userIDUint), input.Amount, input.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with updated wallet balance
	c.JSON(http.StatusOK, gin.H{
		"message": "Withdrawal successful",
		"balance": wallet.Balance,
	})
}
