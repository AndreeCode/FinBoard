package repository

import "finboard/src/core/db/repository"

type RoleRepository struct {
	*repository.CreateRepository
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{
		CreateRepository: repository.NewCreateRepository(),
	}
}
