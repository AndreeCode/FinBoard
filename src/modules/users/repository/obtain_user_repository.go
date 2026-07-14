package repository

import (
	"context"
	"errors"
	"finboard/src/modules/users/domains"
)

var ErrUserNotFound = errors.New("user not found")

func (r *UserRepository) GetUser(ctx context.Context, domain *domains.User) (*domains.User, error) {
	query := `
		SELECT 
			id,
			name,
			lastname,
			email,
			password,
			role_id,
			created_at,
			updated_at,
			deleted_at,
			created_by,
			last_login
		FROM users
		WHERE id = $1
	`
	err := r.DB.QueryRow(
		ctx,
		query,
		domain.Id,
	).Scan(
		&domain.Id,
		&domain.Name,
		&domain.LastName,
		&domain.Email,
		&domain.Password,
		&domain.RoleId,
		&domain.CreatedAt,
		&domain.UpdatedAt,
		&domain.DeletedAt,
		&domain.CreatedBy,
		&domain.LastLogin,
	)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
