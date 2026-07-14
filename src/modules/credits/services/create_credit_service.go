package services

import (
	"context"
	"finboard/src/modules/credits/domains"
)

func (s *CreditService) CreateCredit(
	ctx context.Context,
	credit *domains.Credit,
) (domains.Credit, error) {

	return s.repo.CreateCredit(ctx, credit)
}
