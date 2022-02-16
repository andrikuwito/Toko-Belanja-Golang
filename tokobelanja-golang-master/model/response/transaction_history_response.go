package response

import "time"

type NewTransactionResponse struct {
	Message         string      `json:"message"`
	TransactionBill interface{} `json:"transaction_bill"`
}

type NewTransactionBillResponse struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type TransactionHistoryResponse struct {
	ID         int         `json:"id"`
	ProductID  int         `json:"product_id" valid:"required"`
	UserID     int         `json:"user_id" valid:"required"`
	Quantity   int         `json:"quantity" valid:"required"`
	TotalPrice int         `json:"total_price" valid:"required"`
	Product    interface{} `json:"product"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type UserTransactionHistoryResponse struct {
	ID         int         `json:"id"`
	ProductID  int         `json:"product_id" valid:"required"`
	UserID     int         `json:"user_id" valid:"required"`
	Quantity   int         `json:"quantity" valid:"required"`
	TotalPrice int         `json:"total_price" valid:"required"`
	Product    interface{} `json:"product"`
	User       interface{} `json:"user"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
