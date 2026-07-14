package services

import (
	"context"
	"finboard/src/modules/categories/domains"
)

func (s *CategoryService) DeleteCategory(
	ctx context.Context,
	category *domains.Category,
) error {

	return s.repo.DeleteCategory(ctx, category)
}
