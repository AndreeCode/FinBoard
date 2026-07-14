package repository

import (
	"context"
	"finboard/src/modules/users/domains"
)

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domains.User, error) {
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
		WHERE email = $1 AND deleted_at IS NULL
	`
	var user domains.User
	err := r.DB.QueryRow(
		ctx,
		query,
		email,
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
