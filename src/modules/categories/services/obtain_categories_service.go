package services

import (
	"context"
	"finboard/src/modules/categories/domains"
)

func (s *CategoryService) ObtainCategories(
	ctx context.Context,
	userId string,
) ([]domains.Category, error) {

	return s.repo.GetList(ctx, userId)
}
