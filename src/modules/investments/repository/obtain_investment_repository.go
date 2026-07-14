package repository

import (
	"context"
	"errors"
	"finboard/src/modules/investments/domains"
)

var ErrInvestmentNotFound = errors.New("investment not found")

func (r *InvestmentRepository) GetInvestment(ctx context.Context, domain *domains.Investment) (*domains.Investment, error) {
	query := `
		SELECT 
			id,
			transaction_id,
			expected_gain,
			risk_level,
			status,
			created_at,
			updated_at,
			deleted_at,
			created_by
		FROM investments
		WHERE id = $1
	`
	err := r.DB.QueryRow(
		ctx,
		query,
		domain.Id,
	).Scan(
		&domain.Id,
		&domain.TransactionId,
		&domain.ExpectedGain,
		&domain.RiskLevel,
		&domain.Status,
		&domain.CreatedAt,
		&domain.UpdatedAt,
		&domain.DeletedAt,
		&domain.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
