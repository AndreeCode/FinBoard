package services

import (
	"context"
	"finboard/src/modules/investments/domains"
)

func (s *InvestmentService) ObtainInvestments(
	ctx context.Context,
	userId string,
) ([]domains.Investment, error) {

	return s.repo.GetList(ctx, userId)
}
