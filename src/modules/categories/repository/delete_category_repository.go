package repository

import (
	"context"
	"finboard/src/modules/categories/domains"
)

func (r *CategoryRepository) DeleteCategory(ctx context.Context, domain *domains.Category) error {
	query := `
		UPDATE categories
		SET 
			deleted_at = COALESCE($1, deleted_at)
		WHERE id = $2
	`
	_, err := r.DB.Exec(ctx, query, domain.DeletedAt, domain.Id)
	if err != nil {
		return err
	}
	return nil
}
