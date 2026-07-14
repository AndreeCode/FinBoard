package repository

import (
	"context"
	"finboard/src/core/db/repository"

	"github.com/google/uuid"
)

type InvestmentRepository struct {
	*repository.CreateRepository
}

func NewInvestmentRepository() *InvestmentRepository {
	return &InvestmentRepository{
		CreateRepository: repository.NewCreateRepository(),
	}
}

func (r *InvestmentRepository) GetTransactionUserId(ctx context.Context, transactionId uuid.UUID) (uuid.UUID, error) {
	var userId uuid.UUID
	err := r.DB.QueryRow(ctx, "SELECT user_id FROM transactions WHERE id = $1", transactionId).Scan(&userId)
	return userId, err
}
