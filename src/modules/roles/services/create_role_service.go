package services

import (
	"context"
	"finboard/src/modules/roles/domains"
)

func (s *RoleService) CreateRole(
	ctx context.Context,
	role *domains.Role,
) (domains.Role, error) {

	return s.repo.CreateRole(ctx, role)
}
