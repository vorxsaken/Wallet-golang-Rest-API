package validations

// DepositInput struct to receive deposit requests
type DepositInput struct {
	Amount      float64 `json:"amount" binding:"required,min=0"`
	Description string  `json:"description"`
}

// WithdrawInput struct to receive withdrawal requests
type WithdrawInput struct {
	Amount      float64 `json:"amount" binding:"required,min=0"`
	Description string  `json:"description"`
}

// CreateWalletInput struct to receive wallet creation requests
type CreateWalletInput struct {
	Currency string `json:"currency" binding:"required"`
}

// UpdateWalletInput struct to receive wallet update requests
type UpdateWalletInput struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
