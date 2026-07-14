package services

import (
	"context"
	"finboard/src/modules/transactions/domains"
)

func (s *TransactionService) ObtainTransactions(
	ctx context.Context,
	categoryId, userId string,
) ([]domains.Transaction, error) {

	return s.repo.GetList(ctx, categoryId, userId)
}
