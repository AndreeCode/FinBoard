package services

import (
	"context"
	"finboard/src/modules/users/domains"
)

func (s *UserService) ObtainUser(
	ctx context.Context,
	user *domains.User,
) (*domains.User, error) {
	return s.repo.GetUser(ctx, user)
}
