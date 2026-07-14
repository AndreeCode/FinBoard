package routes

import (
	auth "finboard/src/modules/auth"
	categories "finboard/src/modules/categories"
	credits "finboard/src/modules/credits"
	dashboard "finboard/src/modules/dashboard"
	investments "finboard/src/modules/investments"
	permissions "finboard/src/modules/permissions"
	role_permissions "finboard/src/modules/role_permissions"
	routes "finboard/src/modules/roles"
	transactions "finboard/src/modules/transactions"
	users "finboard/src/modules/users"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App) {

	routes.RegisterRoleRoutes(app)
	users.RegisterUserRoutes(app)
	categories.RegisterCategoryRoutes(app)
	transactions.RegisterTransactionRoutes(app)
	investments.RegisterInvestmentRoutes(app)
	credits.RegisterCreditRoutes(app)
	permissions.RegisterPermissionRoutes(app)
	role_permissions.RegisterRolePermissionRoutes(app)
	auth.RegisterAuthRoutes(app)
	dashboard.RegisterDashboardRoutes(app)
}
