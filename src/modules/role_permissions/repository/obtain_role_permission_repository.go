package repository

import (
	"context"
	"errors"
	"finboard/src/modules/role_permissions/domains"
)

var ErrRolePermissionNotFound = errors.New("role permission not found")

func (r *RolePermissionRepository) GetRolePermission(ctx context.Context, domain *domains.RolePermission) (*domains.RolePermission, error) {
	query := `
		SELECT
			id,
			role_id,
			permission_id,
			created_at,
			updated_at,
			deleted_at,
			created_by
		FROM role_permissions
		WHERE id = $1
	`
	err := r.DB.QueryRow(
		ctx,
		query,
		domain.Id,
	).Scan(
		&domain.Id,
		&domain.RoleId,
		&domain.PermissionId,
		&domain.CreatedAt,
		&domain.UpdatedAt,
		&domain.DeletedAt,
		&domain.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
