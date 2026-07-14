package services

import (
	"context"
	"errors"
	"finboard/src/modules/interfaces"

	"github.com/google/uuid"
)

var ErrUnauthorized = errors.New("unauthorized")

type InvestmentService struct {
	repo interfaces.InvestmentRepositoryInterface
}

func NewInvestmentService(repo interfaces.InvestmentRepositoryInterface) *InvestmentService {
	return &InvestmentService{repo: repo}
}

func (s *InvestmentService) CheckOwnership(ctx context.Context, transactionId uuid.UUID, userId string) error {
	txUserId, err := s.repo.GetTransactionUserId(ctx, transactionId)
	if err != nil {
		return err
	}
	if txUserId.String() != userId {
		return ErrUnauthorized
	}
	return nil
}
