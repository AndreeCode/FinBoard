package services

import (
	"context"
	"finboard/src/modules/permissions/domains"
)

func (s *PermissionService) ObtainPermissions(
	ctx context.Context,
) ([]domains.Permission, error) {

	return s.repo.GetList(ctx)
}
