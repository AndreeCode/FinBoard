package domains

type CreateTransactionRequest struct {
	CategoryId      *string `json:"category_id"`
	Amount          float64 `json:"amount"`
	Type            string  `json:"type"`
	TransactionDate string  `json:"transaction_date"`
	ReceivedDate    *string `json:"received_date"`
	DueDate         *string `json:"due_date"`
	Canceled        *bool   `json:"canceled"`
	Description     string  `json:"description"`
}
type UpdateTransactionRequest struct {
	CategoryId      *string  `json:"category_id"`
	Amount          *float64 `json:"amount"`
	Type            *string  `json:"type"`
	TransactionDate *string  `json:"transaction_date"`
	ReceivedDate    *string  `json:"received_date"`
	DueDate         *string  `json:"due_date"`
	Canceled        *bool    `json:"canceled"`
	Description     *string  `json:"description"`
}
