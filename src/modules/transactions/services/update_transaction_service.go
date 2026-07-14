package services

import (
	"context"
	"finboard/src/modules/transactions/domains"
)

func (s *TransactionService) UpdateTransaction(
	ctx context.Context,
	transaction *domains.Transaction,
) (*domains.Transaction, error) {

	return s.repo.Update(ctx, transaction)
}
