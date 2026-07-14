package services

import (
	"context"
	"finboard/src/mocks"
	rolePermissionsDomains "finboard/src/modules/role_permissions/domains"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRolePermissionService_ObtainRolePermissions_Success(t *testing.T) {
	repo := mocks.NewRolePermissionRepositoryMock()
	now := time.Now()
	rpId := uuid.New()
	roleId := uuid.New()
	permId := uuid.New()
	rp := &rolePermissionsDomains.RolePermission{
		Id:           rpId,
		RoleId:       roleId,
		PermissionId: permId,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	repo.RolePerms[rpId.String()] = rp

	service := NewRolePermissionService(repo)

	rolePermissions, err := service.ObtainRolePermissions(context.Background())

	assert.NoError(t, err)
	assert.Len(t, rolePermissions, 1)
	assert.Equal(t, roleId, rolePermissions[0].RoleId)
}

func TestRolePermissionService_ObtainRolePermissions_Empty(t *testing.T) {
	repo := mocks.NewRolePermissionRepositoryMock()
	service := NewRolePermissionService(repo)

	rolePermissions, err := service.ObtainRolePermissions(context.Background())

	assert.NoError(t, err)
	assert.Len(t, rolePermissions, 0)
}

func TestRolePermissionService_ObtainRolePermission_Success(t *testing.T) {
	repo := mocks.NewRolePermissionRepositoryMock()
	now := time.Now()
	rpId := uuid.New()
	roleId := uuid.New()
	permId := uuid.New()
	rp := &rolePermissionsDomains.RolePermission{
		Id:           rpId,
		RoleId:       roleId,
		PermissionId: permId,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	repo.RolePerms[rpId.String()] = rp

	service := NewRolePermissionService(repo)
	rpToFind := &rolePermissionsDomains.RolePermission{Id: rpId}

	foundRP, err := service.ObtainRolePermission(context.Background(), rpToFind)

	assert.NoError(t, err)
	assert.NotNil(t, foundRP)
	assert.Equal(t, roleId, foundRP.RoleId)
}

func TestRolePermissionService_ObtainRolePermission_NotFound(t *testing.T) {
	repo := mocks.NewRolePermissionRepositoryMock()
	service := NewRolePermissionService(repo)
	rpToFind := &rolePermissionsDomains.RolePermission{Id: uuid.New()}

	foundRP, err := service.ObtainRolePermission(context.Background(), rpToFind)

	assert.Error(t, err)
	assert.Nil(t, foundRP)
}

func TestRolePermissionService_CreateRolePermission_Success(t *testing.T) {
	repo := mocks.NewRolePermissionRepositoryMock()
	service := NewRolePermissionService(repo)
	roleId := uuid.New()
	permId := uuid.New()
	rp := &rolePermissionsDomains.RolePermission{
		RoleId:       roleId,
		PermissionId: permId,
	}

	createdRP, err := service.CreateRolePermission(context.Background(), rp)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, createdRP.Id)
	assert.Equal(t, roleId, createdRP.RoleId)
}

func TestRolePermissionService_DeleteRolePermission_Success(t *testing.T) {
	repo := mocks.NewRolePermissionRepositoryMock()
	now := time.Now()
	rpId := uuid.New()
	roleId := uuid.New()
	permId := uuid.New()
	rp := &rolePermissionsDomains.RolePermission{
		Id:           rpId,
		RoleId:       roleId,
		PermissionId: permId,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	repo.RolePerms[rpId.String()] = rp

	service := NewRolePermissionService(repo)

	err := service.DeleteRolePermission(context.Background(), rp)

	assert.NoError(t, err)
	assert.Len(t, repo.RolePerms, 0)
}

func TestRolePermissionService_NewRolePermissionService(t *testing.T) {
	repo := mocks.NewRolePermissionRepositoryMock()
	service := NewRolePermissionService(repo)

	assert.NotNil(t, service)
}
