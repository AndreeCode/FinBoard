package services

import (
	"context"
	"finboard/src/modules/investments/domains"
)

func (s *InvestmentService) CreateInvestment(
	ctx context.Context,
	investment *domains.Investment,
) (domains.Investment, error) {

	return s.repo.CreateInvestment(ctx, investment)
}
