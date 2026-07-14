package services

import (
	"context"
	"finboard/src/modules/role_permissions/domains"
)

func (s *RolePermissionService) ObtainRolePermissions(
	ctx context.Context,
) ([]domains.RolePermission, error) {

	return s.repo.GetList(ctx)
}
