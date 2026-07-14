package repository

import (
	"context"
	"finboard/src/modules/users/domains"

	"github.com/google/uuid"
)

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*domains.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

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
		WHERE id = $1 AND deleted_at IS NULL
	`
	var user domains.User
	err = r.DB.QueryRow(
		ctx,
		query,
		uid,
	).Scan(
		&user.Id,
		&user.Name,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.RoleId,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.CreatedBy,
		&user.LastLogin,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
