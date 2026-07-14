package repository

import (
	"context"
	"finboard/src/modules/roles/domains"
)

func (r *RoleRepository) DeleteRole(ctx context.Context, domain *domains.Role) error {
	query := `
		UPDATE roles
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
