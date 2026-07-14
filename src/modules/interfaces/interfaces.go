package interfaces

import (
	"context"
	categoriesDomains "finboard/src/modules/categories/domains"
	creditsDomains "finboard/src/modules/credits/domains"
	dashboardDomains "finboard/src/modules/dashboard/domains"
	investmentsDomains "finboard/src/modules/investments/domains"
	permissionsDomains "finboard/src/modules/permissions/domains"
	rolePermissionsDomains "finboard/src/modules/role_permissions/domains"
	rolesDomains "finboard/src/modules/roles/domains"
	transactionsDomains "finboard/src/modules/transactions/domains"
	usersDomains "finboard/src/modules/users/domains"

	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*usersDomains.User, error)
	GetUser(ctx context.Context, user *usersDomains.User) (*usersDomains.User, error)
	GetList(ctx context.Context) ([]usersDomains.User, error)
	CreateUser(ctx context.Context, user *usersDomains.User) (usersDomains.User, error)
	Update(ctx context.Context, user *usersDomains.User) (*usersDomains.User, error)
	DeleteUser(ctx context.Context, user *usersDomains.User) error
	UpdateLastLogin(ctx context.Context, user *usersDomains.User) error
}

type CategoryRepositoryInterface interface {
	GetList(ctx context.Context, userId string) ([]categoriesDomains.Category, error)
	GetCategory(ctx context.Context, domain *categoriesDomains.Category) (*categoriesDomains.Category, error)
	CreateCategory(ctx context.Context, category *categoriesDomains.Category) (categoriesDomains.Category, error)
	Update(ctx context.Context, domain *categoriesDomains.Category) (*categoriesDomains.Category, error)
	DeleteCategory(ctx context.Context, domain *categoriesDomains.Category) error
}

type TransactionRepositoryInterface interface {
	GetList(ctx context.Context, categoryId, userId string) ([]transactionsDomains.Transaction, error)
	GetTransaction(ctx context.Context, domain *transactionsDomains.Transaction) (*transactionsDomains.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *transactionsDomains.Transaction) (transactionsDomains.Transaction, error)
	Update(ctx context.Context, domain *transactionsDomains.Transaction) (*transactionsDomains.Transaction, error)
	DeleteTransaction(ctx context.Context, domain *transactionsDomains.Transaction) error
}

type CreditRepositoryInterface interface {
	ObtainCredits(ctx context.Context, userId string) ([]creditsDomains.Credit, error)
	ObtainCredit(ctx context.Context, id uuid.UUID) (creditsDomains.Credit, error)
	CreateCredit(ctx context.Context, credit *creditsDomains.Credit) (creditsDomains.Credit, error)
	UpdateCredit(ctx context.Context, id uuid.UUID, credit *creditsDomains.Credit) (creditsDomains.Credit, error)
	DeleteCredit(ctx context.Context, id uuid.UUID) error
}

type InvestmentRepositoryInterface interface {
	GetList(ctx context.Context, userId string) ([]investmentsDomains.Investment, error)
	GetInvestment(ctx context.Context, domain *investmentsDomains.Investment) (*investmentsDomains.Investment, error)
	CreateInvestment(ctx context.Context, investment *investmentsDomains.Investment) (investmentsDomains.Investment, error)
	Update(ctx context.Context, domain *investmentsDomains.Investment) (*investmentsDomains.Investment, error)
	DeleteInvestment(ctx context.Context, domain *investmentsDomains.Investment) error
	GetTransactionUserId(ctx context.Context, transactionId uuid.UUID) (uuid.UUID, error)
}

type RoleRepositoryInterface interface {
	GetList(ctx context.Context) ([]rolesDomains.Role, error)
	GetRole(ctx context.Context, domain *rolesDomains.Role) (*rolesDomains.Role, error)
	CreateRole(ctx context.Context, role *rolesDomains.Role) (rolesDomains.Role, error)
	Update(ctx context.Context, domain *rolesDomains.Role) (*rolesDomains.Role, error)
	DeleteRole(ctx context.Context, domain *rolesDomains.Role) error
}

type PermissionRepositoryInterface interface {
	GetList(ctx context.Context) ([]permissionsDomains.Permission, error)
	GetPermission(ctx context.Context, domain *permissionsDomains.Permission) (*permissionsDomains.Permission, error)
	CreatePermission(ctx context.Context, permission *permissionsDomains.Permission) (permissionsDomains.Permission, error)
	Update(ctx context.Context, domain *permissionsDomains.Permission) (*permissionsDomains.Permission, error)
	DeletePermission(ctx context.Context, domain *permissionsDomains.Permission) error
}

type RolePermissionRepositoryInterface interface {
	GetList(ctx context.Context) ([]rolePermissionsDomains.RolePermission, error)
	GetRolePermission(ctx context.Context, domain *rolePermissionsDomains.RolePermission) (*rolePermissionsDomains.RolePermission, error)
	CreateRolePermission(ctx context.Context, rolePermission *rolePermissionsDomains.RolePermission) (rolePermissionsDomains.RolePermission, error)
	DeleteRolePermission(ctx context.Context, domain *rolePermissionsDomains.RolePermission) error
}

type DashboardRepositoryInterface interface {
	GetTransactionTotals(ctx context.Context, userId string) (*dashboardDomains.DashboardSummary, error)
	GetTrendsByPeriod(ctx context.Context, userId string, period string) ([]dashboardDomains.PeriodData, error)
	GetExpensesByCategory(ctx context.Context, userId string) ([]dashboardDomains.CategoryExpense, error)
	GetDailyAverages(ctx context.Context, userId string) (*dashboardDomains.DailyAverage, error)
	GetPeriodComparison(ctx context.Context, userId string) ([]dashboardDomains.PeriodData, error)
}
