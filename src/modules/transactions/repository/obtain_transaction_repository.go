package repository

import (
	"context"
	"errors"
	"finboard/src/modules/transactions/domains"
)

var ErrTransactionNotFound = errors.New("transaction not found")

func (r *TransactionRepository) GetTransaction(ctx context.Context, domain *domains.Transaction) (*domains.Transaction, error) {
	query := `
		SELECT 
			id,
			user_id,
			category_id,
			amount,
			type,
			transaction_date,
			received_date,
			due_date,
			canceled,
			description,
			created_at,
			updated_at,
			deleted_at,
			created_by
		FROM transactions
		WHERE id = $1
	`
	err := r.DB.QueryRow(
		ctx,
		query,
		domain.Id,
	).Scan(
		&domain.Id,
		&domain.UserId,
		&domain.CategoryId,
		&domain.Amount,
		&domain.Type,
		&domain.TransactionDate,
		&domain.ReceivedDate,
		&domain.DueDate,
		&domain.Canceled,
		&domain.Description,
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
