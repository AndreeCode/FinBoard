package repository

import "finboard/src/core/db/repository"

type RolePermissionRepository struct {
	*repository.CreateRepository
}

func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{
		CreateRepository: repository.NewCreateRepository(),
	}
}
