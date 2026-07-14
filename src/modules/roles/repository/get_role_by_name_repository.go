package repository

import (
	"context"
	"finboard/src/modules/roles/domains"
)

func (r *RoleRepository) GetRoleByName(ctx context.Context, name string) (*domains.Role, error) {
	query := `
		SELECT 
			id,
			name,
			description,
			created_at,
			updated_at,
			deleted_at,
			created_by
		FROM roles
		WHERE name = $1 AND deleted_at IS NULL
	`
	var role domains.Role
	err := r.DB.QueryRow(
		ctx,
		query,
		name,
	).Scan(
		&role.Id,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
		&role.DeletedAt,
		&role.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return &role, nil
}
