package services

import (
	"finboard/src/modules/interfaces"
)

type UserService struct {
	repo interfaces.UserRepositoryInterface
}

func NewUserService(repo interfaces.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}
