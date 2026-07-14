package domains

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id              uuid.UUID  `json:"id"`
	UserId          uuid.UUID  `json:"user_id"`
	CategoryId      *uuid.UUID `json:"category_id"`
	Amount          float64    `json:"amount"`
	Type            string     `json:"type"`
	TransactionDate time.Time  `json:"transaction_date"`
	ReceivedDate    *time.Time `json:"received_date"`
	DueDate         *time.Time `json:"due_date"`
	Canceled        bool       `json:"canceled"`
	Description     string     `json:"description"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	CreatedBy       uuid.UUID  `json:"created_by"`
}
