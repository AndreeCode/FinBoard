package repository

import (
	"context"
	"errors"
	"finboard/src/modules/permissions/domains"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

var ErrPermissionAlreadyExists = errors.New("permission already exists")

func (r *PermissionRepository) CreatePermission(
	ctx context.Context,
	permission *domains.Permission,
) (domains.Permission, error) {

	permission.Id = uuid.New()

	query := `
		INSERT INTO permissions (
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
		permission.Id,
		permission.Name,
		permission.Description,
		permission.CreatedBy,
	).Scan(
		&permission.CreatedAt,
		&permission.UpdatedAt,
	)

	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {

			switch pgErr.Code {

			case "23505":

				switch pgErr.ConstraintName {

				case "permissions_name_key":
					return domains.Permission{}, ErrPermissionAlreadyExists
				}
			}
		}

		return domains.Permission{}, err
	}

	return *permission, nil
}
