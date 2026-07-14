package services

import (
	"finboard/src/modules/interfaces"
)

type RoleService struct {
	repo interfaces.RoleRepositoryInterface
}

func NewRoleService(repo interfaces.RoleRepositoryInterface) *RoleService {
	return &RoleService{repo: repo}
}
