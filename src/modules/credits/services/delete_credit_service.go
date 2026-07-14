package services

import (
	"context"

	"github.com/google/uuid"
)

func (s *CreditService) DeleteCredit(
	ctx context.Context,
	id uuid.UUID,
) error {

	return s.repo.DeleteCredit(ctx, id)
}
