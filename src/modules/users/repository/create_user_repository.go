package repository

import (
	"context"
	"errors"
	"finboard/src/modules/users/domains"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrUserAlreadyExists = errors.New("user email already exists")
	ErrRoleNotFound      = errors.New("role not found")
)

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user *domains.User,
) (domains.User, error) {

	user.Id = uuid.New()

	query := `
		INSERT INTO users (
			id,
			name,
			lastname,
			email,
			password,
			role_id,
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING created_at, updated_at
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		user.Id,
		user.Name,
		user.LastName,
		user.Email,
		user.Password,
		user.RoleId,
		user.CreatedBy,
	).Scan(
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {

			switch pgErr.Code {

			case "23505":

				switch pgErr.ConstraintName {

				case "users_email_key":
					return domains.User{}, ErrUserAlreadyExists
				}

			case "23503":

				switch pgErr.ConstraintName {

				case "users_role_id_fkey":
					return domains.User{}, ErrRoleNotFound
				}
			}
		}

		return domains.User{}, err
	}

	return *user, nil
}
