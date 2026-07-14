package repository

import (
	"context"
	"finboard/src/modules/role_permissions/domains"
)

func (r *RolePermissionRepository) GetPermissionsByRole(ctx context.Context, roleID string) ([]domains.RolePermission, error) {
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
		WHERE role_id = $1 AND deleted_at IS NULL
	`
	rows, err := r.DB.Query(ctx, query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolePermissions []domains.RolePermission
	for rows.Next() {
		var rp domains.RolePermission
		err := rows.Scan(
			&rp.Id,
			&rp.RoleId,
			&rp.PermissionId,
			&rp.CreatedAt,
			&rp.UpdatedAt,
			&rp.DeletedAt,
			&rp.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, rp)
	}

	return rolePermissions, nil
}
