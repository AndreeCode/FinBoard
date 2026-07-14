package repository

import (
	"context"
	"finboard/src/modules/credits/domains"
	"time"

	"github.com/google/uuid"
)

func (r *CreditRepository) UpdateCredit(
	ctx context.Context,
	id uuid.UUID,
	credit *domains.Credit,
) (domains.Credit, error) {

	query := `
		UPDATE credits SET
			person_name = COALESCE($1, person_name),
			amount = COALESCE($2, amount),
			interest_rate = COALESCE($3, interest_rate),
			is_creditor = COALESCE($4, is_creditor),
			is_secure = COALESCE($5, is_secure),
			due_date = COALESCE($6, due_date),
			status = COALESCE($7, status),
			updated_at = $8
		WHERE id = $9 AND deleted_at IS NULL
		RETURNING created_at, updated_at
	`

	credit.UpdatedAt = func() *time.Time { t := time.Now(); return &t }()

	err := r.DB.QueryRow(
		ctx,
		query,
		credit.PersonName,
		credit.Amount,
		credit.InterestRate,
		credit.IsCreditor,
		credit.IsSecure,
		credit.DueDate,
		credit.Status,
		credit.UpdatedAt,
		id,
	).Scan(&credit.CreatedAt, &credit.UpdatedAt)

	if err != nil {
		return domains.Credit{}, err
	}

	return *credit, nil
}
