package services

import (
	"context"
	"finboard/src/modules/categories/domains"
)

func (s *CategoryService) ObtainCategory(
	ctx context.Context,
	category *domains.Category,
) (*domains.Category, error) {
	return s.repo.GetCategory(ctx, category)
}
