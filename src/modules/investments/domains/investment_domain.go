package domains

import (
	"time"

	"github.com/google/uuid"
)

type Investment struct {
	Id            uuid.UUID  `json:"id"`
	TransactionId uuid.UUID  `json:"transaction_id"`
	ExpectedGain  float64    `json:"expected_gain"`
	RiskLevel     string     `json:"risk_level"`
	Status        string     `json:"status"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	CreatedBy     uuid.UUID  `json:"created_by"`
}
