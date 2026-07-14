package services

import (
	"context"
	"finboard/src/modules/role_permissions/domains"
)

func (s *RolePermissionService) DeleteRolePermission(
	ctx context.Context,
	rolePermission *domains.RolePermission,
) error {

	return s.repo.DeleteRolePermission(ctx, rolePermission)
}
