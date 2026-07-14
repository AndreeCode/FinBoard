package repository

import (
	"context"
	"finboard/src/modules/users/domains"
	"time"
)

func (r *UserRepository) DeleteUser(ctx context.Context, domain *domains.User) error {
	query := `
		UPDATE users
		SET
			deleted_at = $1
		WHERE id = $2
	`
	now := time.Now()
	_, err := r.DB.Exec(ctx, query, &now, domain.Id)
	if err != nil {
		return err
	}
	return nil
}
