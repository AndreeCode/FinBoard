package services

import (
	"context"
	"errors"
	"finboard/src/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDashboardService_GetDashboard_Success(t *testing.T) {
	repo := mocks.NewDashboardRepositoryMock()
	service := NewDashboardService(repo)
	userId := uuid.New().String()

	dashboard, err := service.GetDashboard(context.Background(), userId, "monthly")

	assert.NoError(t, err)
	assert.NotNil(t, dashboard)
	assert.Equal(t, 1000.0, dashboard.Summary.TotalIncome)
	assert.Equal(t, 500.0, dashboard.Summary.TotalExpenses)
}

func TestDashboardService_GetDashboard_TotalsError(t *testing.T) {
	repo := mocks.NewDashboardRepositoryMock()
	repo.GetTransactionTotalsErr = errors.New("database error")
	service := NewDashboardService(repo)
	userId := uuid.New().String()

	dashboard, err := service.GetDashboard(context.Background(), userId, "monthly")

	assert.Error(t, err)
	assert.Nil(t, dashboard)
}

func TestDashboardService_GetDashboard_TrendsError(t *testing.T) {
	repo := mocks.NewDashboardRepositoryMock()
	repo.GetTrendsByPeriodErr = errors.New("database error")
	service := NewDashboardService(repo)
	userId := uuid.New().String()

	dashboard, err := service.GetDashboard(context.Background(), userId, "monthly")

	assert.NoError(t, err)
	assert.NotNil(t, dashboard)
	assert.Empty(t, dashboard.Trends)
}

func TestDashboardService_GetDashboard_ExpensesByCategoryError(t *testing.T) {
	repo := mocks.NewDashboardRepositoryMock()
	repo.GetExpensesByCategoryErr = errors.New("database error")
	service := NewDashboardService(repo)
	userId := uuid.New().String()

	dashboard, err := service.GetDashboard(context.Background(), userId, "monthly")

	assert.NoError(t, err)
	assert.NotNil(t, dashboard)
	assert.Empty(t, dashboard.ByCategory)
}

func TestDashboardService_GetDashboard_DailyAveragesError(t *testing.T) {
	repo := mocks.NewDashboardRepositoryMock()
	repo.GetDailyAveragesErr = errors.New("database error")
	service := NewDashboardService(repo)
	userId := uuid.New().String()

	dashboard, err := service.GetDashboard(context.Background(), userId, "monthly")

	assert.NoError(t, err)
	assert.NotNil(t, dashboard)
}

func TestDashboardService_GetDashboard_PeriodComparisonError(t *testing.T) {
	repo := mocks.NewDashboardRepositoryMock()
	repo.GetPeriodComparisonErr = errors.New("database error")
	service := NewDashboardService(repo)
	userId := uuid.New().String()

	dashboard, err := service.GetDashboard(context.Background(), userId, "monthly")

	assert.NoError(t, err)
	assert.NotNil(t, dashboard)
	assert.Empty(t, dashboard.PeriodComparison)
}

func TestDashboardService_NewDashboardService(t *testing.T) {
	repo := mocks.NewDashboardRepositoryMock()
	service := NewDashboardService(repo)

	assert.NotNil(t, service)
}
