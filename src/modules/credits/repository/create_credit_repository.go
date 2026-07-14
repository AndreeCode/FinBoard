package repository

import (
	"context"
	"finboard/src/modules/credits/domains"

	"github.com/google/uuid"
)

func (r *CreditRepository) CreateCredit(
	ctx context.Context,
	credit *domains.Credit,
) (domains.Credit, error) {

	credit.Id = uuid.New()

	query := `
		INSERT INTO credits (
			id,
			user_id,
			person_name,
			amount,
			interest_rate,
			is_creditor,
			is_secure,
			due_date,
			status,
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING created_at, updated_at
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		credit.Id,
		credit.UserId,
		credit.PersonName,
		credit.Amount,
		credit.InterestRate,
		credit.IsCreditor,
		credit.IsSecure,
		credit.DueDate,
		credit.Status,
		credit.CreatedBy,
	).Scan(
		&credit.CreatedAt,
		&credit.UpdatedAt,
	)

	if err != nil {
		return domains.Credit{}, err
	}

	return *credit, nil
}
