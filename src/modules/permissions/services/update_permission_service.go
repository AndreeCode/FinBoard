package services

import (
	"context"
	"finboard/src/modules/permissions/domains"
)

func (s *PermissionService) UpdatePermission(
	ctx context.Context,
	permission *domains.Permission,
) (*domains.Permission, error) {

	return s.repo.Update(ctx, permission)
}
