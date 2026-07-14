package repository

import (
	"context"
	"finboard/src/modules/users/domains"
)

func (r *UserRepository) GetList(ctx context.Context) ([]domains.User, error) {
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
	`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domains.User
	for rows.Next() {
		var user domains.User
		err := rows.Scan(
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
		users = append(users, user)
	}

	return users, nil
}
