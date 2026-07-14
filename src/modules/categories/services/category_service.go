package services

import (
	"finboard/src/modules/interfaces"
)

type CategoryService struct {
	repo interfaces.CategoryRepositoryInterface
}

func NewCategoryService(repo interfaces.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{repo: repo}
}
