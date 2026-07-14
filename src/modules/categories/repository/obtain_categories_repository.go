package repository

import (
	"context"
	"finboard/src/modules/categories/domains"
	"fmt"
)

func (r *CategoryRepository) GetList(ctx context.Context, userId string) ([]domains.Category, error) {
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
		WHERE deleted_at IS NULL
	`
	var args []interface{}
	argIndex := 1

	if userId != "" {
		query += fmt.Sprintf(` AND (user_id = $%d OR user_id IS NULL)`, argIndex)
		args = append(args, userId)
		argIndex++
	}

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domains.Category
	for rows.Next() {
		var category domains.Category
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Description,
			&category.ParentId,
			&category.UserId,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.DeletedAt,
			&category.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
