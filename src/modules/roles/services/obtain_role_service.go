package services

import (
	"context"
	"finboard/src/modules/roles/domains"
)

func (s *RoleService) ObtainRole(
	ctx context.Context,
	role *domains.Role,
) (*domains.Role, error) {
	return s.repo.GetRole(ctx, role)
}
