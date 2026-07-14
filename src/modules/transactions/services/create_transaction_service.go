package services

import (
	"context"
	"finboard/src/modules/transactions/domains"
)

func (s *TransactionService) CreateTransaction(
	ctx context.Context,
	transaction *domains.Transaction,
) (domains.Transaction, error) {

	return s.repo.CreateTransaction(ctx, transaction)
}
