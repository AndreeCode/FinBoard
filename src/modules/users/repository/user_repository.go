package repository

import "finboard/src/core/db/repository"

type UserRepository struct {
	*repository.CreateRepository
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		CreateRepository: repository.NewCreateRepository(),
	}
}
