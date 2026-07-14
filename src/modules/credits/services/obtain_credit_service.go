package services

import (
	"context"
	"finboard/src/modules/credits/domains"

	"github.com/google/uuid"
)

func (s *CreditService) ObtainCredit(
	ctx context.Context,
	id uuid.UUID,
) (domains.Credit, error) {

	return s.repo.ObtainCredit(ctx, id)
}
