package repository

import (
	"context"
	"finboard/src/modules/permissions/domains"

	"github.com/google/uuid"
)

func (r *PermissionRepository) GetPermissionByID(ctx context.Context, id string) (*domains.Permission, error) {
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
		FROM permissions
		WHERE id = $1 AND deleted_at IS NULL
	`
	var perm domains.Permission
	err = r.DB.QueryRow(
		ctx,
		query,
		uid,
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
