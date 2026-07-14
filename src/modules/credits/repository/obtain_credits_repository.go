package repository

import (
	"context"
	"finboard/src/modules/credits/domains"
	"fmt"
)

func (r *CreditRepository) ObtainCredits(
	ctx context.Context,
	userId string,
) ([]domains.Credit, error) {
	query := `
		SELECT
			id, user_id, person_name, amount, interest_rate,
			is_creditor, is_secure, due_date, status,
			created_at, updated_at, deleted_at, created_by
		FROM credits
		WHERE deleted_at IS NULL
	`
	var args []interface{}
	argIndex := 1

	if userId != "" {
		query += fmt.Sprintf(` AND user_id = $%d`, argIndex)
		args = append(args, userId)
		argIndex++
	}

	query += ` ORDER BY created_at DESC`

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var credits []domains.Credit
	for rows.Next() {
		var credit domains.Credit
		err := rows.Scan(
			&credit.Id,
			&credit.UserId,
			&credit.PersonName,
			&credit.Amount,
			&credit.InterestRate,
			&credit.IsCreditor,
			&credit.IsSecure,
			&credit.DueDate,
			&credit.Status,
			&credit.CreatedAt,
			&credit.UpdatedAt,
			&credit.DeletedAt,
			&credit.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		credits = append(credits, credit)
	}

	return credits, nil
}
