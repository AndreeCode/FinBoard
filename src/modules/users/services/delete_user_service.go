package services

import (
	"context"
	"finboard/src/modules/users/domains"
)

func (s *UserService) DeleteUser(
	ctx context.Context,
	user *domains.User,
) error {

	return s.repo.DeleteUser(ctx, user)
}
