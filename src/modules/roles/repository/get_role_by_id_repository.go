package repository

import (
	"context"
	"finboard/src/modules/roles/domains"

	"github.com/google/uuid"
)

func (r *RoleRepository) GetRoleByID(ctx context.Context, id string) (*domains.Role, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

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
		WHERE id = $1 AND deleted_at IS NULL
	`
	var role domains.Role
	err = r.DB.QueryRow(
		ctx,
		query,
		uid,
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
