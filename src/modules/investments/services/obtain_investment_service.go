package services

import (
	"context"
	"finboard/src/modules/investments/domains"
)

func (s *InvestmentService) ObtainInvestment(
	ctx context.Context,
	investment *domains.Investment,
) (*domains.Investment, error) {
	return s.repo.GetInvestment(ctx, investment)
}
