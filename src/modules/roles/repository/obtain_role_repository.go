package repository

import (
	"context"
	"finboard/src/modules/roles/domains"
)

func (r *RoleRepository) GetRole(ctx context.Context, domain *domains.Role) (*domains.Role, error) {
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
		WHERE id = $1
	`
	err := r.DB.QueryRow(
		ctx,
		query,
		domain.Id,
	).Scan(
		&domain.Id,
		&domain.Name,
		&domain.Description,
		&domain.CreatedAt,
		&domain.UpdatedAt,
		&domain.DeletedAt,
		&domain.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return domain, nil

}
