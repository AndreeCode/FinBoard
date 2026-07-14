package repository

import (
	"context"
	"finboard/src/core/db/repository"
	"finboard/src/modules/dashboard/domains"
	"fmt"
	"strings"
	"time"
)

type DashboardRepository struct {
	*repository.CreateRepository
}

func NewDashboardRepository() *DashboardRepository {
	return &DashboardRepository{
		CreateRepository: repository.NewCreateRepository(),
	}
}

func (r *DashboardRepository) GetTransactionTotals(ctx context.Context, userId string) (*domains.DashboardSummary, error) {
	var summary domains.DashboardSummary

	query := `
		SELECT
			COALESCE(SUM(CASE WHEN type = 'income' AND deleted_at IS NULL THEN amount ELSE 0 END), 0) as total_income,
			COALESCE(SUM(CASE WHEN type = 'expense' AND deleted_at IS NULL THEN amount ELSE 0 END), 0) as total_expenses,
			COALESCE(SUM(CASE WHEN type = 'investment' AND deleted_at IS NULL THEN amount ELSE 0 END), 0) as total_investments,
			COUNT(CASE WHEN type = 'income' AND deleted_at IS NULL THEN 1 END) as income_count,
			COUNT(CASE WHEN type = 'expense' AND deleted_at IS NULL THEN 1 END) as expense_count
		FROM transactions
		WHERE deleted_at IS NULL AND user_id = $1
	`

	var incomeCount, expenseCount int
	err := r.DB.QueryRow(ctx, query, userId).Scan(
		&summary.TotalIncome,
		&summary.TotalExpenses,
		&summary.TotalInvestments,
		&incomeCount,
		&expenseCount,
	)
	if err != nil {
		return nil, err
	}

	creditQuery := `
		SELECT
			COALESCE(SUM(CASE WHEN is_creditor = true THEN amount ELSE 0 END), 0) as you_are_owed,
			COALESCE(SUM(CASE WHEN is_creditor = false THEN amount ELSE 0 END), 0) as you_owe
		FROM credits
		WHERE deleted_at IS NULL AND user_id = $1
	`
	err = r.DB.QueryRow(ctx, creditQuery, userId).Scan(&summary.YouAreOwed, &summary.YouOwe)
	if err != nil {
		summary.YouAreOwed = 0
		summary.YouOwe = 0
	}

	summary.TotalCredits = summary.YouAreOwed + summary.YouOwe
	summary.Balance = summary.TotalIncome - summary.TotalExpenses - summary.TotalInvestments + summary.YouAreOwed - summary.YouOwe
	summary.Savings = summary.TotalIncome - summary.TotalExpenses - summary.TotalInvestments

	if incomeCount > 0 {
		summary.AverageIncome = summary.TotalIncome / float64(incomeCount)
	}
	if expenseCount > 0 {
		summary.AverageExpenses = summary.TotalExpenses / float64(expenseCount)
	}

	if summary.TotalIncome > 0 {
		summary.SavingsRate = (summary.Savings / summary.TotalIncome) * 100
	}

	return &summary, nil
}

func (r *DashboardRepository) GetTrendsByPeriod(ctx context.Context, userId string, period string) ([]domains.PeriodData, error) {
	var dateFormat string
	var dateTrunc string

	switch period {
	case "daily":
		dateFormat = "YYYY-MM-DD"
		dateTrunc = "day"
	case "weekly":
		dateFormat = "IYYY-IW"
		dateTrunc = "week"
	case "monthly":
		dateFormat = "YYYY-MM"
		dateTrunc = "month"
	case "quarterly":
		dateFormat = "YYYY-Q"
		dateTrunc = "quarter"
	case "semiannual":
		dateFormat = "YYYY-Half"
		dateTrunc = "quarter"
	case "annual":
		dateFormat = "YYYY"
		dateTrunc = "year"
	default:
		dateFormat = "YYYY-MM"
		dateTrunc = "month"
	}

	query := fmt.Sprintf(`
		SELECT
			TO_CHAR(DATE_TRUNC('%s', transaction_date), '%s') as period,
			COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) as income,
			COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) as expenses,
			COALESCE(SUM(CASE WHEN type = 'investment' THEN amount ELSE 0 END), 0) as investments,
			COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE -amount END), 0) as net_flow
		FROM transactions
		WHERE deleted_at IS NULL AND user_id = $1
		GROUP BY TO_CHAR(DATE_TRUNC('%s', transaction_date), '%s')
		ORDER BY period DESC
		LIMIT 24
	`, dateTrunc, dateFormat, dateTrunc, dateFormat)

	rows, err := r.DB.Query(ctx, query, userId)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return []domains.PeriodData{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	var trends []domains.PeriodData
	for rows.Next() {
		var t domains.PeriodData
		if err := rows.Scan(&t.Period, &t.Income, &t.Expenses, &t.Investments, &t.NetFlow); err != nil {
			return nil, err
		}
		trends = append(trends, t)
	}

	return trends, nil
}

func (r *DashboardRepository) GetExpensesByCategory(ctx context.Context, userId string) ([]domains.CategoryExpense, error) {
	query := `
		SELECT
			COALESCE(t.category_id, '') as category_id,
			COALESCE(c.name, 'Sin categoría') as category_name,
			COALESCE(SUM(t.amount), 0) as total,
			COUNT(*) as count
		FROM transactions t
		LEFT JOIN categories c ON t.category_id = c.id
		WHERE t.type = 'expense' AND t.deleted_at IS NULL AND t.user_id = $1
		GROUP BY t.category_id, c.name
		ORDER BY total DESC
	`

	rows, err := r.DB.Query(ctx, query, userId)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return []domains.CategoryExpense{}, nil
		}
		return nil, err
	}
	defer rows.Close()

	var expenses []domains.CategoryExpense
	var grandTotal float64

	for rows.Next() {
		var e domains.CategoryExpense
		if err := rows.Scan(&e.CategoryId, &e.CategoryName, &e.Total, &e.Count); err != nil {
			return nil, err
		}
		grandTotal += e.Total
		expenses = append(expenses, e)
	}

	for i := range expenses {
		if grandTotal > 0 {
			expenses[i].Percentage = (expenses[i].Total / grandTotal) * 100
		}
	}

	return expenses, nil
}

func (r *DashboardRepository) GetDailyAverages(ctx context.Context, userId string) (*domains.DailyAverage, error) {
	query := `
		WITH date_range AS (
			SELECT
				COALESCE(MIN(transaction_date), CURRENT_DATE) as min_date,
				COALESCE(MAX(transaction_date), CURRENT_DATE) as max_date,
				GREATEST(EXTRACT(DAY FROM MAX(transaction_date) - MIN(transaction_date)) + 1, 1) as total_days,
				GREATEST(CEIL(EXTRACT(DAY FROM MAX(transaction_date) - MIN(transaction_date)) / 7.0), 1) as total_weeks,
				GREATEST(EXTRACT(MONTH FROM AGE(MAX(transaction_date), MIN(transaction_date))) + EXTRACT(YEAR FROM AGE(MAX(transaction_date), MIN(transaction_date))) * 12 + 1, 1) as total_months
			FROM transactions
			WHERE deleted_at IS NULL AND user_id = $1
		),
		totals AS (
			SELECT
				COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) as total_income,
				COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) as total_expenses
			FROM transactions
			WHERE deleted_at IS NULL AND user_id = $1
		)
		SELECT
			dr.total_days,
			dr.total_weeks,
			dr.total_months,
			t.total_income,
			t.total_expenses
		FROM date_range dr, totals t
	`

	var days, weeks, months int
	var totalIncome, totalExpenses float64

	err := r.DB.QueryRow(ctx, query, userId).Scan(&days, &weeks, &months, &totalIncome, &totalExpenses)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return &domains.DailyAverage{}, nil
		}
		return nil, err
	}

	return &domains.DailyAverage{
		DailyIncome:      totalIncome / float64(days),
		DailyExpenses:    totalExpenses / float64(days),
		WeeklyIncome:     totalIncome / float64(weeks),
		WeeklyExpenses:   totalExpenses / float64(weeks),
		MonthlyIncome:   totalIncome / float64(months),
		MonthlyExpenses: totalExpenses / float64(months),
	}, nil
}

func (r *DashboardRepository) GetPeriodComparison(ctx context.Context, userId string) ([]domains.PeriodData, error) {
	now := time.Now()

	currentMonthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	lastMonthStart := currentMonthStart.AddDate(0, -1, 0)
	lastMonthEnd := currentMonthStart.AddDate(0, 0, -1)

	lastMonthStartStr := lastMonthStart.Format("2006-01-02")
	lastMonthEndStr := lastMonthEnd.Format("2006-01-02")
	currentMonthStartStr := currentMonthStart.Format("2006-01-02")
	nowStr := now.Format("2006-01-02")

	query := `
		SELECT
			COALESCE(SUM(CASE WHEN type = 'income' AND transaction_date >= $2 AND transaction_date <= $3 THEN amount ELSE 0 END), 0) as current_income,
			COALESCE(SUM(CASE WHEN type = 'expense' AND transaction_date >= $2 AND transaction_date <= $3 THEN amount ELSE 0 END), 0) as current_expenses,
			COALESCE(SUM(CASE WHEN type = 'investment' AND transaction_date >= $2 AND transaction_date <= $3 THEN amount ELSE 0 END), 0) as current_investments,
			COALESCE(SUM(CASE WHEN type = 'income' AND transaction_date >= $4 AND transaction_date <= $5 THEN amount ELSE 0 END), 0) as prev_income,
			COALESCE(SUM(CASE WHEN type = 'expense' AND transaction_date >= $4 AND transaction_date <= $5 THEN amount ELSE 0 END), 0) as prev_expenses
		FROM transactions
		WHERE deleted_at IS NULL AND user_id = $1
	`

	var current, prev domains.PeriodData

	err := r.DB.QueryRow(ctx, query,
		userId,
		currentMonthStartStr, nowStr,
		lastMonthStartStr, lastMonthEndStr,
	).Scan(
		&current.Income, &current.Expenses, &current.Investments,
		&prev.Income, &prev.Expenses,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return []domains.PeriodData{{Period: "current_month"}, {Period: "last_month"}}, nil
		}
		return nil, err
	}

	current.Period = "current_month"
	prev.Period = "last_month"
	current.NetFlow = current.Income - current.Expenses
	prev.NetFlow = prev.Income - prev.Expenses

	return []domains.PeriodData{current, prev}, nil
}
