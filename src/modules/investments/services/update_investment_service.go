package services

import (
	"context"
	"finboard/src/modules/investments/domains"
)

func (s *InvestmentService) UpdateInvestment(
	ctx context.Context,
	investment *domains.Investment,
) (*domains.Investment, error) {

	return s.repo.Update(ctx, investment)
}
