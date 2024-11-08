package validations

// TransactionInput struct to receive transaction requests
type TransactionInput struct {
	WalletID          uint    `json:"wallet_id" binding:"required"`
	ProductID         uint    `json:"product_id" binding:"required"`
	TransactionTypeID uint    `json:"transaction_type_id" binding:"required"`
	Amount            float64 `json:"amount" binding:"required,gt=0"`
	ProductAmount     int     `json:"product_amount" binding:"required,gt=0"`
	Description       string  `json:"description"`
}
