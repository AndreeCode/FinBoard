package repository

import (
	"context"
	"errors"
	"finboard/src/modules/categories/domains"
)

var ErrCategoryNotFound = errors.New("category not found")

func (r *CategoryRepository) GetCategory(ctx context.Context, domain *domains.Category) (*domains.Category, error) {
	query := `
		SELECT 
			id,
			name,
			description,
			parent_id,
			user_id,
			created_at,
			updated_at,
			deleted_at,
			created_by
		FROM categories
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
		&domain.ParentId,
		&domain.UserId,
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
