package domains

import (
	"time"

	"github.com/google/uuid"
)

type Credit struct {
	Id           uuid.UUID  `json:"id"`
	UserId       uuid.UUID  `json:"user_id"`
	PersonName   string     `json:"person_name"`
	Amount       float64    `json:"amount"`
	InterestRate float64    `json:"interest_rate"`
	IsCreditor   bool       `json:"is_creditor"`
	IsSecure     bool       `json:"is_secure"`
	DueDate      *time.Time `json:"due_date"`
	Status       string     `json:"status"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	CreatedBy    uuid.UUID  `json:"created_by"`
}
