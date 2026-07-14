package repository

import (
	"context"
	"errors"
	"finboard/src/modules/role_permissions/domains"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrRolePermissionAlreadyExists = errors.New("role permission already exists")
	ErrRoleNotFoundForPermission   = errors.New("role not found")
	ErrPermissionNotFoundForRole   = errors.New("permission not found")
)

func (r *RolePermissionRepository) CreateRolePermission(
	ctx context.Context,
	rolePermission *domains.RolePermission,
) (domains.RolePermission, error) {

	rolePermission.Id = uuid.New()

	query := `
		INSERT INTO role_permissions (
			id,
			role_id,
			permission_id,
			created_by
		)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at, updated_at
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		rolePermission.Id,
		rolePermission.RoleId,
		rolePermission.PermissionId,
		rolePermission.CreatedBy,
	).Scan(
		&rolePermission.CreatedAt,
		&rolePermission.UpdatedAt,
	)

	if err != nil {

		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {

			switch pgErr.Code {

			case "23505":
				return domains.RolePermission{}, ErrRolePermissionAlreadyExists

			case "23503":

				switch pgErr.ConstraintName {

				case "role_permissions_role_id_fkey":
					return domains.RolePermission{}, ErrRoleNotFoundForPermission

				case "role_permissions_permission_id_fkey":
					return domains.RolePermission{}, ErrPermissionNotFoundForRole
				}
			}
		}

		return domains.RolePermission{}, err
	}

	return *rolePermission, nil
}
