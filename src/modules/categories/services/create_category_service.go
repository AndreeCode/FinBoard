package services

import (
	"context"
	"finboard/src/modules/categories/domains"
)

func (s *CategoryService) CreateCategory(
	ctx context.Context,
	category *domains.Category,
) (domains.Category, error) {

	return s.repo.CreateCategory(ctx, category)
}
