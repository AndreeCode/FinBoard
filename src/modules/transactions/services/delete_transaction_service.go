package services

import (
	"context"
	"finboard/src/modules/transactions/domains"
)

func (s *TransactionService) DeleteTransaction(
	ctx context.Context,
	transaction *domains.Transaction,
) error {

	return s.repo.DeleteTransaction(ctx, transaction)
}
