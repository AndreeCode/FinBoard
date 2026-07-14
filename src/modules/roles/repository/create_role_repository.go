package repository

import (
	"context"
	"errors"
	"finboard/src/modules/roles/domains"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrRoleAlreadyExists     = errors.New("role already exists")
	ErrRoleDescriptionExists = errors.New("role description already exists")
)

func (r *RoleRepository) CreateRole(
	ctx context.Context,
	role *domains.Role,
) (domains.Role, error) {

	role.Id = uuid.New()

	query := `
		INSERT INTO roles (
			id,
			name,
			description,
			created_by
		)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at, updated_at
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		role.Id,
		role.Name,
		role.Description,
		role.CreatedBy,
	).Scan(
		&role.CreatedAt,
		&role.UpdatedAt,
	)

	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {

			switch pgErr.Code {

			case "23505":

				switch pgErr.ConstraintName {

				case "roles_name_key":
					return domains.Role{}, ErrRoleAlreadyExists

				case "roles_description_key":
					return domains.Role{}, ErrRoleDescriptionExists
				}
			}
		}

		return domains.Role{}, err
	}

	return *role, nil
}
