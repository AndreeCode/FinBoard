package services

import (
	"finboard/src/modules/interfaces"
)

type TransactionService struct {
	repo interfaces.TransactionRepositoryInterface
}

func NewTransactionService(repo interfaces.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{repo: repo}
}
