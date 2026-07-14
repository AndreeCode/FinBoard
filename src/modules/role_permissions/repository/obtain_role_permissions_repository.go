package repository

import (
	"context"
	"finboard/src/modules/role_permissions/domains"
)

func (r *RolePermissionRepository) GetList(ctx context.Context) ([]domains.RolePermission, error) {
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
	`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []domains.RolePermission
	for rows.Next() {
		var rolePermission domains.RolePermission
		err := rows.Scan(
			&rolePermission.Id,
			&rolePermission.RoleId,
			&rolePermission.PermissionId,
			&rolePermission.CreatedAt,
			&rolePermission.UpdatedAt,
			&rolePermission.DeletedAt,
			&rolePermission.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	return rolePermissions, nil
}
