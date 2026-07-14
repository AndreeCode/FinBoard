package repository

import (
	"context"
	"finboard/src/modules/transactions/domains"

	"github.com/google/uuid"
)

func (r *TransactionRepository) CreateTransaction(
	ctx context.Context,
	transaction *domains.Transaction,
) (domains.Transaction, error) {

	transaction.Id = uuid.New()

	query := `
		INSERT INTO transactions (
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
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING created_at, updated_at
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		transaction.Id,
		transaction.UserId,
		transaction.CategoryId,
		transaction.Amount,
		transaction.Type,
		transaction.TransactionDate,
		transaction.ReceivedDate,
		transaction.DueDate,
		transaction.Canceled,
		transaction.Description,
		transaction.CreatedBy,
	).Scan(
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		return domains.Transaction{}, err
	}

	return *transaction, nil
}
