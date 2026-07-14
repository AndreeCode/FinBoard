package services

import (
	"context"
	"finboard/src/modules/permissions/domains"
)

func (s *PermissionService) CreatePermission(
	ctx context.Context,
	permission *domains.Permission,
) (domains.Permission, error) {

	return s.repo.CreatePermission(ctx, permission)
}
