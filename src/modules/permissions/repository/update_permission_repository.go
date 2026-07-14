package repository

import (
	"context"
	"finboard/src/modules/permissions/domains"
)

func (r *PermissionRepository) Update(ctx context.Context, domain *domains.Permission) (*domains.Permission, error) {
	query := `
		UPDATE permissions
		SET
			name = COALESCE($1, name),
			description = COALESCE($2, description),
			updated_at = COALESCE($3, updated_at)
		WHERE id = $4
	`
	_, err := r.DB.Exec(ctx, query, domain.Name, domain.Description, domain.UpdatedAt, domain.Id)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
