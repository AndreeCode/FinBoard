package repository

import "finboard/src/core/db/repository"

type CreditRepository struct {
	*repository.CreateRepository
}

func NewCreditRepository() *CreditRepository {
	return &CreditRepository{
		CreateRepository: repository.NewCreateRepository(),
	}
}
