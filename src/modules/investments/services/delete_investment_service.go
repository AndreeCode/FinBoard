package services

import (
	"context"
	"finboard/src/modules/investments/domains"
)

func (s *InvestmentService) DeleteInvestment(
	ctx context.Context,
	investment *domains.Investment,
) error {

	return s.repo.DeleteInvestment(ctx, investment)
}
