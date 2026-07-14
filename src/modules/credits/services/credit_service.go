package services

import (
	"finboard/src/modules/interfaces"
)

type CreditService struct {
	repo interfaces.CreditRepositoryInterface
}

func NewCreditService(repo interfaces.CreditRepositoryInterface) *CreditService {
	return &CreditService{repo: repo}
}
