package repository

import (
	"context"
	"finboard/src/modules/investments/domains"

	"github.com/google/uuid"
)

func (r *InvestmentRepository) CreateInvestment(
	ctx context.Context,
	investment *domains.Investment,
) (domains.Investment, error) {

	investment.Id = uuid.New()

	query := `
		INSERT INTO investments (
			id,
			transaction_id,
			expected_gain,
			risk_level,
			status,
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING created_at, updated_at
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		investment.Id,
		investment.TransactionId,
		investment.ExpectedGain,
		investment.RiskLevel,
		investment.Status,
		investment.CreatedBy,
	).Scan(
		&investment.CreatedAt,
		&investment.UpdatedAt,
	)

	if err != nil {
		return domains.Investment{}, err
	}

	return *investment, nil
}
