package repository

import (
	"context"
	"finboard/src/modules/roles/domains"
)

func (r *RoleRepository) GetList(ctx context.Context) ([]domains.Role, error) {
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
	`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domains.Role
	for rows.Next() {
		var role domains.Role
		err := rows.Scan(
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
		roles = append(roles, role)
	}

	return roles, nil
}
