package services

import (
	"context"
	"finboard/src/modules/users/domains"
)

func (s *UserService) UpdateUser(
	ctx context.Context,
	user *domains.User,
) (*domains.User, error) {

	return s.repo.Update(ctx, user)
}
