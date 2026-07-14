package services

import (
	"context"
	"finboard/src/mocks"
	rolesDomains "finboard/src/modules/roles/domains"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRoleService_ObtainRoles_Success(t *testing.T) {
	repo := mocks.NewRoleRepositoryMock()
	now := time.Now()
	roleId := uuid.New()
	role := &rolesDomains.Role{
		Id:          roleId,
		Name:        "admin",
		Description: "Administrator role",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repo.Roles[roleId.String()] = role

	service := NewRoleService(repo)

	roles, err := service.ObtainRoles(context.Background())

	assert.NoError(t, err)
	assert.Len(t, roles, 1)
	assert.Equal(t, "admin", roles[0].Name)
}

func TestRoleService_ObtainRoles_Empty(t *testing.T) {
	repo := mocks.NewRoleRepositoryMock()
	service := NewRoleService(repo)

	roles, err := service.ObtainRoles(context.Background())

	assert.NoError(t, err)
	assert.Len(t, roles, 0)
}

func TestRoleService_ObtainRole_Success(t *testing.T) {
	repo := mocks.NewRoleRepositoryMock()
	now := time.Now()
	roleId := uuid.New()
	role := &rolesDomains.Role{
		Id:          roleId,
		Name:        "admin",
		Description: "Administrator role",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repo.Roles[roleId.String()] = role

	service := NewRoleService(repo)
	roleToFind := &rolesDomains.Role{Id: roleId}

	foundRole, err := service.ObtainRole(context.Background(), roleToFind)

	assert.NoError(t, err)
	assert.NotNil(t, foundRole)
	assert.Equal(t, "admin", foundRole.Name)
}

func TestRoleService_ObtainRole_NotFound(t *testing.T) {
	repo := mocks.NewRoleRepositoryMock()
	service := NewRoleService(repo)
	roleToFind := &rolesDomains.Role{Id: uuid.New()}

	foundRole, err := service.ObtainRole(context.Background(), roleToFind)

	assert.Error(t, err)
	assert.Nil(t, foundRole)
}

func TestRoleService_CreateRole_Success(t *testing.T) {
	repo := mocks.NewRoleRepositoryMock()
	service := NewRoleService(repo)
	role := &rolesDomains.Role{
		Name:        "editor",
		Description: "Editor role",
	}

	createdRole, err := service.CreateRole(context.Background(), role)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, createdRole.Id)
	assert.Equal(t, "editor", createdRole.Name)
}

func TestRoleService_UpdateRole_Success(t *testing.T) {
	repo := mocks.NewRoleRepositoryMock()
	now := time.Now()
	roleId := uuid.New()
	role := &rolesDomains.Role{
		Id:          roleId,
		Name:        "admin",
		Description: "Administrator role",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repo.Roles[roleId.String()] = role

	service := NewRoleService(repo)
	role.Name = "superadmin"

	updatedRole, err := service.UpdateRole(context.Background(), role)

	assert.NoError(t, err)
	assert.NotNil(t, updatedRole)
	assert.Equal(t, "superadmin", updatedRole.Name)
}

func TestRoleService_UpdateRole_NotFound(t *testing.T) {
	repo := mocks.NewRoleRepositoryMock()
	service := NewRoleService(repo)
	role := &rolesDomains.Role{
		Id:   uuid.New(),
		Name: "superadmin",
	}

	updatedRole, err := service.UpdateRole(context.Background(), role)

	assert.Error(t, err)
	assert.Nil(t, updatedRole)
}

func TestRoleService_DeleteRole_Success(t *testing.T) {
	repo := mocks.NewRoleRepositoryMock()
	now := time.Now()
	roleId := uuid.New()
	role := &rolesDomains.Role{
		Id:          roleId,
		Name:        "admin",
		Description: "Administrator role",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repo.Roles[roleId.String()] = role

	service := NewRoleService(repo)

	err := service.DeleteRole(context.Background(), role)

	assert.NoError(t, err)
	assert.Len(t, repo.Roles, 0)
}

func TestRoleService_NewRoleService(t *testing.T) {
	repo := mocks.NewRoleRepositoryMock()
	service := NewRoleService(repo)

	assert.NotNil(t, service)
}
