package repository

import (
	"context"
	"finboard/src/modules/permissions/domains"
)

func (r *PermissionRepository) GetList(ctx context.Context) ([]domains.Permission, error) {
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
	`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []domains.Permission
	for rows.Next() {
		var permission domains.Permission
		err := rows.Scan(
			&permission.Id,
			&permission.Name,
			&permission.Description,
			&permission.CreatedAt,
			&permission.UpdatedAt,
			&permission.DeletedAt,
			&permission.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
