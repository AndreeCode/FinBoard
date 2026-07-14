package repository

import (
	"context"
	"finboard/src/modules/credits/domains"

	"github.com/google/uuid"
)

func (r *CreditRepository) ObtainCredit(
	ctx context.Context,
	id uuid.UUID,
) (domains.Credit, error) {

	query := `
		SELECT
			id, user_id, person_name, amount, interest_rate,
			is_creditor, is_secure, due_date, status,
			created_at, updated_at, deleted_at, created_by
		FROM credits
		WHERE id = $1 AND deleted_at IS NULL
	`

	var credit domains.Credit
	err := r.DB.QueryRow(ctx, query, id).Scan(
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
		return domains.Credit{}, err
	}

	return credit, nil
}
