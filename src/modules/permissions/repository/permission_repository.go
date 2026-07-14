package repository

import "finboard/src/core/db/repository"

type PermissionRepository struct {
	*repository.CreateRepository
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{
		CreateRepository: repository.NewCreateRepository(),
	}
}
