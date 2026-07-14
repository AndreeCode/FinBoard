package services

import (
	"context"
	"finboard/src/modules/role_permissions/domains"
)

func (s *RolePermissionService) ObtainRolePermission(
	ctx context.Context,
	rolePermission *domains.RolePermission,
) (*domains.RolePermission, error) {
	return s.repo.GetRolePermission(ctx, rolePermission)
}
