package repository

import (
	"context"
	"finboard/src/modules/permissions/domains"
)

func (r *PermissionRepository) GetPermissionByName(ctx context.Context, name string) (*domains.Permission, error) {
	query := `
		SELECT
			id,
			name,
			description,
			created_at,
			updated_at,
			deleted_at,
			created_by
		FROM permissions
		WHERE name = $1 AND deleted_at IS NULL
	`
	var perm domains.Permission
	err := r.DB.QueryRow(
		ctx,
		query,
		name,
	).Scan(
		&perm.Id,
		&perm.Name,
		&perm.Description,
		&perm.CreatedAt,
		&perm.UpdatedAt,
		&perm.DeletedAt,
		&perm.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return &perm, nil
}
