package services

import (
	"context"
	"finboard/src/modules/users/domains"
)

func (s *UserService) ObtainUsers(
	ctx context.Context,
) ([]domains.User, error) {

	return s.repo.GetList(ctx)
}
