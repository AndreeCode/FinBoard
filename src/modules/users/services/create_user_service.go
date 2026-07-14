package services

import (
	"context"
	"finboard/src/core/utils"
	"finboard/src/modules/users/domains"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	user *domains.User,
) (domains.User, error) {

	hashedPassword, err := utils.HashPassword(user.Password, nil)
	if err != nil {
		return domains.User{}, err
	}
	user.Password = hashedPassword

	return s.repo.CreateUser(ctx, user)
}
