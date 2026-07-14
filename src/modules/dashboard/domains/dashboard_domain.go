package domains

type DashboardSummary struct {
	TotalIncome      float64            `json:"total_income"`
	TotalExpenses    float64            `json:"total_expenses"`
	TotalInvestments float64            `json:"total_investments"`
	TotalCredits     float64            `json:"total_credits"`
	YouOwe           float64            `json:"you_owe"`
	YouAreOwed       float64            `json:"you_are_owed"`
	Balance          float64            `json:"balance"`
	AverageIncome    float64            `json:"average_income"`
	AverageExpenses  float64            `json:"average_expenses"`
	Savings          float64            `json:"savings"`
	SavingsRate      float64            `json:"savings_rate"`
}

type PeriodData struct {
	Period      string  `json:"period"`
	Income      float64 `json:"income"`
	Expenses    float64 `json:"expenses"`
	Investments float64 `json:"investments"`
	Credits     float64 `json:"credits"`
	YouOwe      float64 `json:"you_owe"`
	YouAreOwed  float64 `json:"you_are_owed"`
	NetFlow     float64 `json:"net_flow"`
}

type CategoryExpense struct {
	CategoryId   string  `json:"category_id"`
	CategoryName string  `json:"category_name"`
	Total        float64 `json:"total"`
	Percentage   float64 `json:"percentage"`
	Count        int     `json:"count"`
}

type DailyAverage struct {
	DailyIncome     float64 `json:"daily_income"`
	DailyExpenses   float64 `json:"daily_expenses"`
	WeeklyIncome    float64 `json:"weekly_income"`
	WeeklyExpenses  float64 `json:"weekly_expenses"`
	MonthlyIncome   float64 `json:"monthly_income"`
	MonthlyExpenses float64 `json:"monthly_expenses"`
}

type DashboardResponse struct {
	Summary          DashboardSummary   `json:"summary"`
	Trends           []PeriodData      `json:"trends"`
	ByCategory       []CategoryExpense `json:"by_category"`
	DailyAverage     DailyAverage      `json:"daily_average"`
	PeriodComparison []PeriodData      `json:"period_comparison"`
}
