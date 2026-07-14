package repository

import (
	"context"
	"finboard/src/modules/transactions/domains"
)

func (r *TransactionRepository) Update(ctx context.Context, domain *domains.Transaction) (*domains.Transaction, error) {
	query := `
		UPDATE transactions
		SET 
			category_id = COALESCE($1, category_id),
			amount = COALESCE($2, amount),
			type = COALESCE($3, type),
			transaction_date = COALESCE($4, transaction_date),
			received_date = COALESCE($5, received_date),
			due_date = COALESCE($6, due_date),
			canceled = COALESCE($7, canceled),
			description = COALESCE($8, description),
			updated_at = COALESCE($9, updated_at)
		WHERE id = $10
	`
	_, err := r.DB.Exec(ctx, query,
		domain.CategoryId,
		domain.Amount,
		domain.Type,
		domain.TransactionDate,
		domain.ReceivedDate,
		domain.DueDate,
		domain.Canceled,
		domain.Description,
		domain.UpdatedAt,
		domain.Id,
	)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
