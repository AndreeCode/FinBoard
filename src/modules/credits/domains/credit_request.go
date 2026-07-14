package domains

type CreateCreditRequest struct {
	PersonName   string   `json:"person_name"`
	Amount       float64  `json:"amount"`
	InterestRate float64  `json:"interest_rate"`
	IsCreditor   bool     `json:"is_creditor"`
	IsSecure     bool     `json:"is_secure"`
	DueDate      *string  `json:"due_date"`
	Status       string   `json:"status"`
}

type UpdateCreditRequest struct {
	PersonName   *string  `json:"person_name"`
	Amount       *float64 `json:"amount"`
	InterestRate *float64 `json:"interest_rate"`
	IsCreditor   *bool    `json:"is_creditor"`
	IsSecure     *bool    `json:"is_secure"`
	DueDate      *string  `json:"due_date"`
	Status       *string  `json:"status"`
}
