package repository

import "finboard/src/core/db/repository"

type TransactionRepository struct {
	*repository.CreateRepository
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		CreateRepository: repository.NewCreateRepository(),
	}
}
