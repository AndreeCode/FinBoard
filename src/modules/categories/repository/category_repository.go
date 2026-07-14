package repository

import "finboard/src/core/db/repository"

type CategoryRepository struct {
	*repository.CreateRepository
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		CreateRepository: repository.NewCreateRepository(),
	}
}
