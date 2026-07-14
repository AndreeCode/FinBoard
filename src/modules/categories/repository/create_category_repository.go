package repository

import (
	"context"
	"finboard/src/modules/categories/domains"

	"github.com/google/uuid"
)

func (r *CategoryRepository) CreateCategory(
	ctx context.Context,
	category *domains.Category,
) (domains.Category, error) {

	category.Id = uuid.New()

	query := `
		INSERT INTO categories (
			id,
			name,
			description,
			parent_id,
			user_id,
			created_by
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING created_at, updated_at
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		category.Id,
		category.Name,
		category.Description,
		category.ParentId,
		category.UserId,
		category.CreatedBy,
	).Scan(
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		return domains.Category{}, err
	}

	return *category, nil
}
