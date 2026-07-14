package services

import (
	"context"
	"finboard/src/modules/permissions/domains"
)

func (s *PermissionService) ObtainPermission(
	ctx context.Context,
	permission *domains.Permission,
) (*domains.Permission, error) {
	return s.repo.GetPermission(ctx, permission)
}
