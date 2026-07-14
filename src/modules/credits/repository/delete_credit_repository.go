package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (r *CreditRepository) DeleteCredit(
	ctx context.Context,
	id uuid.UUID,
) error {

	query := `
		UPDATE credits SET deleted_at = $2 WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.DB.Exec(ctx, query, id, time.Now())
	return err
}
