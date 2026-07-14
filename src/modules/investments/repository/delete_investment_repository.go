package repository

import (
	"context"
	"finboard/src/modules/investments/domains"
)

func (r *InvestmentRepository) DeleteInvestment(ctx context.Context, domain *domains.Investment) error {
	query := `
		UPDATE investments
		SET 
			deleted_at = COALESCE($1, deleted_at)
		WHERE id = $2
	`
	_, err := r.DB.Exec(ctx, query, domain.DeletedAt, domain.Id)
	if err != nil {
		return err
	}
	return nil
}
