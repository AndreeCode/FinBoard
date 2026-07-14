package repository

import (
	"context"
	"finboard/src/modules/investments/domains"
)

func (r *InvestmentRepository) Update(ctx context.Context, domain *domains.Investment) (*domains.Investment, error) {
	query := `
		UPDATE investments
		SET 
			expected_gain = COALESCE($1, expected_gain),
			risk_level = COALESCE($2, risk_level),
			status = COALESCE($3, status),
			updated_at = COALESCE($4, updated_at)
		WHERE id = $5
	`
	_, err := r.DB.Exec(ctx, query,
		domain.ExpectedGain,
		domain.RiskLevel,
		domain.Status,
		domain.UpdatedAt,
		domain.Id,
	)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
