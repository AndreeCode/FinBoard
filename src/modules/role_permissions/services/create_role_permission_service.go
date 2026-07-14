package services

import (
	"context"
	"finboard/src/modules/role_permissions/domains"
)

func (s *RolePermissionService) CreateRolePermission(
	ctx context.Context,
	rolePermission *domains.RolePermission,
) (domains.RolePermission, error) {

	return s.repo.CreateRolePermission(ctx, rolePermission)
}
