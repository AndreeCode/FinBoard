package services

import (
	"context"
	"finboard/src/modules/categories/domains"
)

func (s *CategoryService) UpdateCategory(
	ctx context.Context,
	category *domains.Category,
) (*domains.Category, error) {

	return s.repo.Update(ctx, category)
}
