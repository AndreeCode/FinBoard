package repository

import (
	"context"
	"finboard/src/modules/transactions/domains"
	"fmt"
	"strings"
)

func (r *TransactionRepository) GetList(ctx context.Context, categoryId, userId string) ([]domains.Transaction, error) {
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
		WHERE deleted_at IS NULL
	`
	var args []interface{}
	argIndex := 1

	if userId != "" {
		query += fmt.Sprintf(` AND user_id = $%d`, argIndex)
		args = append(args, userId)
		argIndex++
	}

	if categoryId != "" {
		query += fmt.Sprintf(` AND category_id = $%d`, argIndex)
		args = append(args, categoryId)
		argIndex++
	}

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		if strings.Contains(err.Error(), "missing") {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	var transactions []domains.Transaction
	for rows.Next() {
		var transaction domains.Transaction
		err := rows.Scan(
			&transaction.Id,
			&transaction.UserId,
			&transaction.CategoryId,
			&transaction.Amount,
			&transaction.Type,
			&transaction.TransactionDate,
			&transaction.ReceivedDate,
			&transaction.DueDate,
			&transaction.Canceled,
			&transaction.Description,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.DeletedAt,
			&transaction.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
