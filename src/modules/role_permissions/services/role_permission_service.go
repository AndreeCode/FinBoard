package services

import (
	"finboard/src/modules/interfaces"
)

type RolePermissionService struct {
	repo interfaces.RolePermissionRepositoryInterface
}

func NewRolePermissionService(repo interfaces.RolePermissionRepositoryInterface) *RolePermissionService {
	return &RolePermissionService{repo: repo}
}
