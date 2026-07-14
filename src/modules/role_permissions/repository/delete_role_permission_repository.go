package repository

import (
	"context"
	"finboard/src/modules/role_permissions/domains"
)

func (r *RolePermissionRepository) DeleteRolePermission(ctx context.Context, domain *domains.RolePermission) error {
	query := `
		UPDATE role_permissions
		SET
			deleted_at = COALESCE($1, deleted_at)
		WHERE id = $2
	`
	_, err := r.DB.Exec(ctx, query, domain.DeletedAt, domain.Id)
	if err != nil {
		return err
	}
	return nil
}
