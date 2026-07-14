package services

import (
	"context"
	"finboard/src/modules/credits/domains"

	"github.com/google/uuid"
)

func (s *CreditService) UpdateCredit(
	ctx context.Context,
	id uuid.UUID,
	credit *domains.Credit,
) (domains.Credit, error) {

	return s.repo.UpdateCredit(ctx, id, credit)
}
