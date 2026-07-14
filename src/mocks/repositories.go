package mocks

import (
	"context"
	"errors"
	categoriesDomains "finboard/src/modules/categories/domains"
	creditsDomains "finboard/src/modules/credits/domains"
	dashboardDomains "finboard/src/modules/dashboard/domains"
	investmentsDomains "finboard/src/modules/investments/domains"
	permissionsDomains "finboard/src/modules/permissions/domains"
	rolePermissionsDomains "finboard/src/modules/role_permissions/domains"
	rolesDomains "finboard/src/modules/roles/domains"
	transactionsDomains "finboard/src/modules/transactions/domains"
	usersDomains "finboard/src/modules/users/domains"
	"sync"
	"time"

	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")
var ErrCategoryNotFound = errors.New("category not found")
var ErrTransactionNotFound = errors.New("transaction not found")
var ErrCreditNotFound = errors.New("credit not found")

type UserRepositoryMock struct {
	mu    sync.RWMutex
	Users map[string]*usersDomains.User
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{
		Users: make(map[string]*usersDomains.User),
	}
}

func (m *UserRepositoryMock) GetUserByEmail(ctx context.Context, email string) (*usersDomains.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, u := range m.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

func (m *UserRepositoryMock) GetUser(ctx context.Context, user *usersDomains.User) (*usersDomains.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if u, ok := m.Users[user.Id.String()]; ok {
		return u, nil
	}
	return nil, ErrUserNotFound
}

func (m *UserRepositoryMock) GetList(ctx context.Context) ([]usersDomains.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	users := make([]usersDomains.User, 0, len(m.Users))
	for _, u := range m.Users {
		users = append(users, *u)
	}
	return users, nil
}

func (m *UserRepositoryMock) CreateUser(ctx context.Context, user *usersDomains.User) (usersDomains.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if user.Id == uuid.Nil {
		user.Id = uuid.New()
	}
	now := time.Now()
	user.CreatedAt = &now
	user.UpdatedAt = &now
	m.Users[user.Id.String()] = user
	return *user, nil
}

func (m *UserRepositoryMock) Update(ctx context.Context, user *usersDomains.User) (*usersDomains.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.Users[user.Id.String()]; ok {
		now := time.Now()
		user.UpdatedAt = &now
		m.Users[user.Id.String()] = user
		return user, nil
	}
	return nil, ErrUserNotFound
}

func (m *UserRepositoryMock) DeleteUser(ctx context.Context, user *usersDomains.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.Users[user.Id.String()]; ok {
		delete(m.Users, user.Id.String())
		return nil
	}
	return ErrUserNotFound
}

func (m *UserRepositoryMock) UpdateLastLogin(ctx context.Context, user *usersDomains.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	now := time.Now()
	user.LastLogin = &now
	return nil
}

type CategoryRepositoryMock struct {
	mu         sync.RWMutex
	Categories map[string]*categoriesDomains.Category
}

func NewCategoryRepositoryMock() *CategoryRepositoryMock {
	return &CategoryRepositoryMock{
		Categories: make(map[string]*categoriesDomains.Category),
	}
}

func (m *CategoryRepositoryMock) GetList(ctx context.Context, userId string) ([]categoriesDomains.Category, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	cats := make([]categoriesDomains.Category, 0)
	for _, c := range m.Categories {
		if c.UserId != nil && c.UserId.String() == userId {
			cats = append(cats, *c)
		}
	}
	return cats, nil
}

func (m *CategoryRepositoryMock) GetCategory(ctx context.Context, domain *categoriesDomains.Category) (*categoriesDomains.Category, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if c, ok := m.Categories[domain.Id.String()]; ok {
		return c, nil
	}
	return nil, ErrCategoryNotFound
}

func (m *CategoryRepositoryMock) CreateCategory(ctx context.Context, cat *categoriesDomains.Category) (categoriesDomains.Category, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if cat.Id == uuid.Nil {
		cat.Id = uuid.New()
	}
	now := time.Now()
	cat.CreatedAt = &now
	cat.UpdatedAt = &now
	m.Categories[cat.Id.String()] = cat
	return *cat, nil
}

func (m *CategoryRepositoryMock) Update(ctx context.Context, domain *categoriesDomains.Category) (*categoriesDomains.Category, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.Categories[domain.Id.String()]; ok {
		now := time.Now()
		domain.UpdatedAt = &now
		m.Categories[domain.Id.String()] = domain
		return domain, nil
	}
	return nil, nil
}

func (m *CategoryRepositoryMock) DeleteCategory(ctx context.Context, domain *categoriesDomains.Category) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.Categories, domain.Id.String())
	return nil
}

type TransactionRepositoryMock struct {
	mu           sync.RWMutex
	Transactions map[string]*transactionsDomains.Transaction
}

func NewTransactionRepositoryMock() *TransactionRepositoryMock {
	return &TransactionRepositoryMock{
		Transactions: make(map[string]*transactionsDomains.Transaction),
	}
}

func (m *TransactionRepositoryMock) GetList(ctx context.Context, categoryId, userId string) ([]transactionsDomains.Transaction, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	txs := make([]transactionsDomains.Transaction, 0)
	for _, t := range m.Transactions {
		if (userId == "" || t.UserId.String() == userId) && (categoryId == "" || t.CategoryId.String() == categoryId) {
			txs = append(txs, *t)
		}
	}
	return txs, nil
}

func (m *TransactionRepositoryMock) GetTransaction(ctx context.Context, domain *transactionsDomains.Transaction) (*transactionsDomains.Transaction, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if t, ok := m.Transactions[domain.Id.String()]; ok {
		return t, nil
	}
	return nil, ErrTransactionNotFound
}

func (m *TransactionRepositoryMock) CreateTransaction(ctx context.Context, tx *transactionsDomains.Transaction) (transactionsDomains.Transaction, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if tx.Id == uuid.Nil {
		tx.Id = uuid.New()
	}
	now := time.Now()
	tx.CreatedAt = &now
	tx.UpdatedAt = &now
	m.Transactions[tx.Id.String()] = tx
	return *tx, nil
}

func (m *TransactionRepositoryMock) Update(ctx context.Context, domain *transactionsDomains.Transaction) (*transactionsDomains.Transaction, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.Transactions[domain.Id.String()]; ok {
		now := time.Now()
		domain.UpdatedAt = &now
		m.Transactions[domain.Id.String()] = domain
		return domain, nil
	}
	return nil, ErrTransactionNotFound
}

func (m *TransactionRepositoryMock) DeleteTransaction(ctx context.Context, domain *transactionsDomains.Transaction) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.Transactions, domain.Id.String())
	return nil
}

type CreditRepositoryMock struct {
	mu      sync.RWMutex
	Credits map[string]*creditsDomains.Credit
}

func NewCreditRepositoryMock() *CreditRepositoryMock {
	return &CreditRepositoryMock{
		Credits: make(map[string]*creditsDomains.Credit),
	}
}

func (m *CreditRepositoryMock) ObtainCredits(ctx context.Context, userId string) ([]creditsDomains.Credit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	credits := make([]creditsDomains.Credit, 0)
	for _, c := range m.Credits {
		if c.UserId.String() == userId {
			credits = append(credits, *c)
		}
	}
	return credits, nil
}

func (m *CreditRepositoryMock) ObtainCredit(ctx context.Context, id uuid.UUID) (creditsDomains.Credit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if c, ok := m.Credits[id.String()]; ok {
		return *c, nil
	}
	return creditsDomains.Credit{}, ErrCreditNotFound
}

func (m *CreditRepositoryMock) CreateCredit(ctx context.Context, credit *creditsDomains.Credit) (creditsDomains.Credit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if credit.Id == uuid.Nil {
		credit.Id = uuid.New()
	}
	now := time.Now()
	credit.CreatedAt = &now
	credit.UpdatedAt = &now
	m.Credits[credit.Id.String()] = credit
	return *credit, nil
}

func (m *CreditRepositoryMock) UpdateCredit(ctx context.Context, id uuid.UUID, credit *creditsDomains.Credit) (creditsDomains.Credit, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.Credits[id.String()]; ok {
		now := time.Now()
		credit.UpdatedAt = &now
		m.Credits[id.String()] = credit
		return *credit, nil
	}
	return creditsDomains.Credit{}, ErrCreditNotFound
}

func (m *CreditRepositoryMock) DeleteCredit(ctx context.Context, id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.Credits, id.String())
	return nil
}

type InvestmentRepositoryMock struct {
	mu              sync.RWMutex
	Investments     map[string]*investmentsDomains.Investment
	TxUserId        uuid.UUID
	TxUserIdErr     error
}

func NewInvestmentRepositoryMock() *InvestmentRepositoryMock {
	return &InvestmentRepositoryMock{
		Investments: make(map[string]*investmentsDomains.Investment),
	}
}

func (m *InvestmentRepositoryMock) GetList(ctx context.Context, userId string) ([]investmentsDomains.Investment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	investments := make([]investmentsDomains.Investment, 0)
	for _, i := range m.Investments {
		investments = append(investments, *i)
	}
	return investments, nil
}

func (m *InvestmentRepositoryMock) GetInvestment(ctx context.Context, domain *investmentsDomains.Investment) (*investmentsDomains.Investment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if i, ok := m.Investments[domain.Id.String()]; ok {
		return i, nil
	}
	return nil, ErrInvestmentNotFound
}

func (m *InvestmentRepositoryMock) CreateInvestment(ctx context.Context, inv *investmentsDomains.Investment) (investmentsDomains.Investment, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if inv.Id == uuid.Nil {
		inv.Id = uuid.New()
	}
	now := time.Now()
	inv.CreatedAt = &now
	inv.UpdatedAt = &now
	m.Investments[inv.Id.String()] = inv
	return *inv, nil
}

func (m *InvestmentRepositoryMock) Update(ctx context.Context, domain *investmentsDomains.Investment) (*investmentsDomains.Investment, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.Investments[domain.Id.String()]; ok {
		now := time.Now()
		domain.UpdatedAt = &now
		m.Investments[domain.Id.String()] = domain
		return domain, nil
	}
	return nil, ErrInvestmentNotFound
}

func (m *InvestmentRepositoryMock) DeleteInvestment(ctx context.Context, domain *investmentsDomains.Investment) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.Investments, domain.Id.String())
	return nil
}

func (m *InvestmentRepositoryMock) GetTransactionUserId(ctx context.Context, transactionId uuid.UUID) (uuid.UUID, error) {
	if m.TxUserIdErr != nil {
		return uuid.Nil, m.TxUserIdErr
	}
	if m.TxUserId != uuid.Nil {
		return m.TxUserId, nil
	}
	return uuid.Nil, nil
}

var ErrInvestmentNotFound = errors.New("investment not found")

type RoleRepositoryMock struct {
	mu    sync.RWMutex
	Roles map[string]*rolesDomains.Role
}

func NewRoleRepositoryMock() *RoleRepositoryMock {
	return &RoleRepositoryMock{
		Roles: make(map[string]*rolesDomains.Role),
	}
}

func (m *RoleRepositoryMock) GetList(ctx context.Context) ([]rolesDomains.Role, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	roles := make([]rolesDomains.Role, 0, len(m.Roles))
	for _, r := range m.Roles {
		roles = append(roles, *r)
	}
	return roles, nil
}

func (m *RoleRepositoryMock) GetRole(ctx context.Context, domain *rolesDomains.Role) (*rolesDomains.Role, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if r, ok := m.Roles[domain.Id.String()]; ok {
		return r, nil
	}
	return nil, ErrRoleNotFound
}

func (m *RoleRepositoryMock) CreateRole(ctx context.Context, role *rolesDomains.Role) (rolesDomains.Role, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if role.Id == uuid.Nil {
		role.Id = uuid.New()
	}
	now := time.Now()
	role.CreatedAt = &now
	role.UpdatedAt = &now
	m.Roles[role.Id.String()] = role
	return *role, nil
}

func (m *RoleRepositoryMock) Update(ctx context.Context, domain *rolesDomains.Role) (*rolesDomains.Role, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.Roles[domain.Id.String()]; ok {
		now := time.Now()
		domain.UpdatedAt = &now
		m.Roles[domain.Id.String()] = domain
		return domain, nil
	}
	return nil, ErrRoleNotFound
}

func (m *RoleRepositoryMock) DeleteRole(ctx context.Context, domain *rolesDomains.Role) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.Roles, domain.Id.String())
	return nil
}

var ErrRoleNotFound = errors.New("role not found")

type PermissionRepositoryMock struct {
	mu          sync.RWMutex
	Permissions map[string]*permissionsDomains.Permission
}

func NewPermissionRepositoryMock() *PermissionRepositoryMock {
	return &PermissionRepositoryMock{
		Permissions: make(map[string]*permissionsDomains.Permission),
	}
}

func (m *PermissionRepositoryMock) GetList(ctx context.Context) ([]permissionsDomains.Permission, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	perms := make([]permissionsDomains.Permission, 0, len(m.Permissions))
	for _, p := range m.Permissions {
		perms = append(perms, *p)
	}
	return perms, nil
}

func (m *PermissionRepositoryMock) GetPermission(ctx context.Context, domain *permissionsDomains.Permission) (*permissionsDomains.Permission, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if p, ok := m.Permissions[domain.Id.String()]; ok {
		return p, nil
	}
	return nil, ErrPermissionNotFound
}

func (m *PermissionRepositoryMock) CreatePermission(ctx context.Context, perm *permissionsDomains.Permission) (permissionsDomains.Permission, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if perm.Id == uuid.Nil {
		perm.Id = uuid.New()
	}
	now := time.Now()
	perm.CreatedAt = &now
	perm.UpdatedAt = &now
	m.Permissions[perm.Id.String()] = perm
	return *perm, nil
}

func (m *PermissionRepositoryMock) Update(ctx context.Context, domain *permissionsDomains.Permission) (*permissionsDomains.Permission, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.Permissions[domain.Id.String()]; ok {
		now := time.Now()
		domain.UpdatedAt = &now
		m.Permissions[domain.Id.String()] = domain
		return domain, nil
	}
	return nil, ErrPermissionNotFound
}

func (m *PermissionRepositoryMock) DeletePermission(ctx context.Context, domain *permissionsDomains.Permission) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.Permissions, domain.Id.String())
	return nil
}

var ErrPermissionNotFound = errors.New("permission not found")

type RolePermissionRepositoryMock struct {
	mu        sync.RWMutex
	RolePerms map[string]*rolePermissionsDomains.RolePermission
}

func NewRolePermissionRepositoryMock() *RolePermissionRepositoryMock {
	return &RolePermissionRepositoryMock{
		RolePerms: make(map[string]*rolePermissionsDomains.RolePermission),
	}
}

func (m *RolePermissionRepositoryMock) GetList(ctx context.Context) ([]rolePermissionsDomains.RolePermission, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	rps := make([]rolePermissionsDomains.RolePermission, 0, len(m.RolePerms))
	for _, rp := range m.RolePerms {
		rps = append(rps, *rp)
	}
	return rps, nil
}

func (m *RolePermissionRepositoryMock) GetRolePermission(ctx context.Context, domain *rolePermissionsDomains.RolePermission) (*rolePermissionsDomains.RolePermission, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if rp, ok := m.RolePerms[domain.Id.String()]; ok {
		return rp, nil
	}
	return nil, ErrRolePermissionNotFound
}

func (m *RolePermissionRepositoryMock) CreateRolePermission(ctx context.Context, rp *rolePermissionsDomains.RolePermission) (rolePermissionsDomains.RolePermission, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if rp.Id == uuid.Nil {
		rp.Id = uuid.New()
	}
	now := time.Now()
	rp.CreatedAt = &now
	m.RolePerms[rp.Id.String()] = rp
	return *rp, nil
}

func (m *RolePermissionRepositoryMock) DeleteRolePermission(ctx context.Context, domain *rolePermissionsDomains.RolePermission) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.RolePerms, domain.Id.String())
	return nil
}

var ErrRolePermissionNotFound = errors.New("role permission not found")

type DashboardRepositoryMock struct {
	GetTransactionTotalsErr  error
	GetTrendsByPeriodErr    error
	GetExpensesByCategoryErr error
	GetDailyAveragesErr     error
	GetPeriodComparisonErr  error
}

func NewDashboardRepositoryMock() *DashboardRepositoryMock {
	return &DashboardRepositoryMock{}
}

func (m *DashboardRepositoryMock) GetTransactionTotals(ctx context.Context, userId string) (*dashboardDomains.DashboardSummary, error) {
	if m.GetTransactionTotalsErr != nil {
		return nil, m.GetTransactionTotalsErr
	}
	return &dashboardDomains.DashboardSummary{
		TotalIncome:   1000,
		TotalExpenses: 500,
		Balance:       500,
	}, nil
}

func (m *DashboardRepositoryMock) GetTrendsByPeriod(ctx context.Context, userId string, period string) ([]dashboardDomains.PeriodData, error) {
	if m.GetTrendsByPeriodErr != nil {
		return nil, m.GetTrendsByPeriodErr
	}
	return []dashboardDomains.PeriodData{}, nil
}

func (m *DashboardRepositoryMock) GetExpensesByCategory(ctx context.Context, userId string) ([]dashboardDomains.CategoryExpense, error) {
	if m.GetExpensesByCategoryErr != nil {
		return nil, m.GetExpensesByCategoryErr
	}
	return []dashboardDomains.CategoryExpense{}, nil
}

func (m *DashboardRepositoryMock) GetDailyAverages(ctx context.Context, userId string) (*dashboardDomains.DailyAverage, error) {
	if m.GetDailyAveragesErr != nil {
		return nil, m.GetDailyAveragesErr
	}
	return &dashboardDomains.DailyAverage{}, nil
}

func (m *DashboardRepositoryMock) GetPeriodComparison(ctx context.Context, userId string) ([]dashboardDomains.PeriodData, error) {
	if m.GetPeriodComparisonErr != nil {
		return nil, m.GetPeriodComparisonErr
	}
	return []dashboardDomains.PeriodData{}, nil
}
