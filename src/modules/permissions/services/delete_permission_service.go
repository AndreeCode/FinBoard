package services

import (
	"context"
	"finboard/src/modules/permissions/domains"
)

func (s *PermissionService) DeletePermission(
	ctx context.Context,
	permission *domains.Permission,
) error {

	return s.repo.DeletePermission(ctx, permission)
}
