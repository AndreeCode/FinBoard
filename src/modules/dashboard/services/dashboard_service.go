package services

import (
	"context"
	"finboard/src/modules/dashboard/domains"
	"finboard/src/modules/interfaces"
)

type DashboardService struct {
	repo interfaces.DashboardRepositoryInterface
}

func NewDashboardService(repo interfaces.DashboardRepositoryInterface) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetDashboard(ctx context.Context, userId string, period string) (*domains.DashboardResponse, error) {
	summary, err := s.repo.GetTransactionTotals(ctx, userId)
	if err != nil {
		return nil, err
	}

	trends, err := s.repo.GetTrendsByPeriod(ctx, userId, period)
	if err != nil {
		trends = []domains.PeriodData{}
	}

	byCategory, err := s.repo.GetExpensesByCategory(ctx, userId)
	if err != nil {
		byCategory = []domains.CategoryExpense{}
	}

	dailyAverage, err := s.repo.GetDailyAverages(ctx, userId)
	if err != nil {
		dailyAverage = &domains.DailyAverage{}
	}

	periodComparison, err := s.repo.GetPeriodComparison(ctx, userId)
	if err != nil {
		periodComparison = []domains.PeriodData{}
	}

	return &domains.DashboardResponse{
		Summary:         *summary,
		Trends:          trends,
		ByCategory:      byCategory,
		DailyAverage:    *dailyAverage,
		PeriodComparison: periodComparison,
	}, nil
}
