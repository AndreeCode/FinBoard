package repository

import (
	"context"
	"finboard/src/modules/investments/domains"
	"fmt"
)

func (r *InvestmentRepository) GetList(ctx context.Context, userId string) ([]domains.Investment, error) {
	query := `
		SELECT
			i.id,
			i.transaction_id,
			i.expected_gain,
			i.risk_level,
			i.status,
			i.created_at,
			i.updated_at,
			i.deleted_at,
			i.created_by
		FROM investments i
		INNER JOIN transactions t ON i.transaction_id = t.id
		WHERE i.deleted_at IS NULL AND t.deleted_at IS NULL
	`
	var args []interface{}
	argIndex := 1

	if userId != "" {
		query += fmt.Sprintf(` AND t.user_id = $%d`, argIndex)
		args = append(args, userId)
		argIndex++
	}

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var investments []domains.Investment
	for rows.Next() {
		var investment domains.Investment
		err := rows.Scan(
			&investment.Id,
			&investment.TransactionId,
			&investment.ExpectedGain,
			&investment.RiskLevel,
			&investment.Status,
			&investment.CreatedAt,
			&investment.UpdatedAt,
			&investment.DeletedAt,
			&investment.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		investments = append(investments, investment)
	}

	return investments, nil
}
