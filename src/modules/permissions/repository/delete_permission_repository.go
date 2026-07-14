package repository

import (
	"context"
	"finboard/src/modules/permissions/domains"
)

func (r *PermissionRepository) DeletePermission(ctx context.Context, domain *domains.Permission) error {
	query := `
		UPDATE permissions
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
