package services

import (
	"context"
	"finboard/src/modules/roles/domains"
)

func (s *RoleService) UpdateRole(
	ctx context.Context,
	role *domains.Role,
) (*domains.Role, error) {

	return s.repo.Update(ctx, role)
}
