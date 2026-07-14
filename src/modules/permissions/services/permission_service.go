package services

import (
	"finboard/src/modules/interfaces"
)

type PermissionService struct {
	repo interfaces.PermissionRepositoryInterface
}

func NewPermissionService(repo interfaces.PermissionRepositoryInterface) *PermissionService {
	return &PermissionService{repo: repo}
}
