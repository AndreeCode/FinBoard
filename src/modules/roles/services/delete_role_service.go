package services

import (
	"context"
	"finboard/src/modules/roles/domains"
)

func (s *RoleService) DeleteRole(
	ctx context.Context,
	role *domains.Role,
) error {

	return s.repo.DeleteRole(ctx, role)
}
