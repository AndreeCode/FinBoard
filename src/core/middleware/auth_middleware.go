package middleware

import (
	"finboard/src/core/utils"
	hndlers "finboard/src/modules/auth/handlers"
	permRepo "finboard/src/modules/permissions/repository"
	rolePermRepo "finboard/src/modules/role_permissions/repository"
	roleRepo "finboard/src/modules/roles/repository"

	"github.com/gofiber/fiber/v3"
)

type AuthMiddleware struct {
	rolePermRepo *rolePermRepo.RolePermissionRepository
	roleRepo     *roleRepo.RoleRepository
	permRepo     *permRepo.PermissionRepository
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		rolePermRepo: rolePermRepo.NewRolePermissionRepository(),
		roleRepo:     roleRepo.NewRoleRepository(),
		permRepo:     permRepo.NewPermissionRepository(),
	}
}
func (m *AuthMiddleware) RequireAuth() fiber.Handler {
	return func(c fiber.Ctx) error {

		token := hndlers.GetAuthToken(c)
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": fiber.StatusUnauthorized,
				"msg":    "unauthorized",
				"data":   nil,
			})
		}

		claims, err := utils.ValidateAccessToken(token, nil)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": fiber.StatusUnauthorized,
				"msg":    "invalid token",
				"data":   nil,
			})
		}

		c.Locals("user_id", claims.UserID.String())
		c.Locals("name", claims.Name)
		c.Locals("last_name", claims.LastName)
		c.Locals("email", claims.Email)
		c.Locals("role_id", claims.RoleID.String())

		return c.Next()
	}
}

func (m *AuthMiddleware) RequirePermission(permission string) fiber.Handler {
	return func(c fiber.Ctx) error {
		roleID := c.Locals("role_id")
		if roleID == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": fiber.StatusUnauthorized,
				"msg":    "unauthorized",
				"data":   nil,
			})
		}

		roleIDStr := roleID.(string)

		role, err := m.roleRepo.GetRoleByID(c.Context(), roleIDStr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": fiber.StatusInternalServerError,
				"msg":    "error checking role",
				"data":   nil,
			})
		}

		if role.Name == "admin" {
			return c.Next()
		}

		perms, err := m.rolePermRepo.GetPermissionsByRole(c.Context(), roleIDStr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": fiber.StatusInternalServerError,
				"msg":    "error checking permissions",
				"data":   nil,
			})
		}

		hasPermission := false
		for _, rp := range perms {
			perm, err := m.permRepo.GetPermissionByID(c.Context(), rp.PermissionId.String())
			if err == nil && perm.Name == permission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status": fiber.StatusForbidden,
				"msg":    "permission denied",
				"data":   nil,
			})
		}

		return c.Next()
	}
}

func (m *AuthMiddleware) IsAdmin(c fiber.Ctx) bool {
	roleID := c.Locals("role_id")
	if roleID == nil {
		return false
	}

	role, err := m.roleRepo.GetRoleByID(c.Context(), roleID.(string))
	if err != nil {
		return false
	}

	return role.Name == "admin"
}

func (m *AuthMiddleware) CanAccessResource(c fiber.Ctx, resourceUserId string) bool {
	currentUserID := c.Locals("user_id")
	if currentUserID == nil {
		return false
	}

	if currentUserID.(string) == resourceUserId {
		return true
	}

	if m.IsAdmin(c) {
		return true
	}

	return false
}
