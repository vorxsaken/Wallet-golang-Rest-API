package controllers

import (
	"fahmi-wallet/services"
	"fahmi-wallet/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct{}

var transactionService = services.TransactionService{}

func (s *TransactionController) InitiateTransaction(c *gin.Context) {
	var input validations.TransactionInput

	// Bind and validate the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Call service to handle the transaction
	transaction, err := transactionService.InitiateTransaction(
		input.WalletID,
		input.ProductID,
		input.TransactionTypeID,
		input.Amount,
		input.Description,
		input.ProductAmount,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}
