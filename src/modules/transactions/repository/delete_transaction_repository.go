package repository

import (
	"context"
	"finboard/src/modules/transactions/domains"
)

func (r *TransactionRepository) DeleteTransaction(ctx context.Context, domain *domains.Transaction) error {
	query := `
		UPDATE transactions
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
