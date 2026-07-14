package services

import (
	"context"
	"finboard/src/modules/transactions/domains"
)

func (s *TransactionService) ObtainTransaction(
	ctx context.Context,
	transaction *domains.Transaction,
) (*domains.Transaction, error) {
	return s.repo.GetTransaction(ctx, transaction)
}
