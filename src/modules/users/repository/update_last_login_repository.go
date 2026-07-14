package repository

import (
	"context"
	"finboard/src/modules/users/domains"
	"time"
)

func (r *UserRepository) UpdateLastLogin(ctx context.Context, user *domains.User) error {
	now := time.Now()
	query := `
		UPDATE users
		SET last_login = $1
		WHERE id = $2
	`
	_, err := r.DB.Exec(ctx, query, now, user.Id)
	if err != nil {
		return err
	}
	user.LastLogin = &now
	return nil
}
