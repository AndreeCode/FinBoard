package services

import (
	"context"
	"finboard/src/modules/roles/domains"
)

func (s *RoleService) ObtainRoles(
	ctx context.Context,
) ([]domains.Role, error) {

	return s.repo.GetList(ctx)
}
