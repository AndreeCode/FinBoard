package services

import (
	"context"
	"finboard/src/modules/credits/domains"
)

func (s *CreditService) ObtainCredits(
	ctx context.Context,
	userId string,
) ([]domains.Credit, error) {

	return s.repo.ObtainCredits(ctx, userId)
}
