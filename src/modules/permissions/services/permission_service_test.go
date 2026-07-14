package services

import (
	"context"
	"finboard/src/mocks"
	permissionsDomains "finboard/src/modules/permissions/domains"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPermissionService_ObtainPermissions_Success(t *testing.T) {
	repo := mocks.NewPermissionRepositoryMock()
	now := time.Now()
	permId := uuid.New()
	perm := &permissionsDomains.Permission{
		Id:          permId,
		Name:        "read",
		Description: "Read permission",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repo.Permissions[permId.String()] = perm

	service := NewPermissionService(repo)

	permissions, err := service.ObtainPermissions(context.Background())

	assert.NoError(t, err)
	assert.Len(t, permissions, 1)
	assert.Equal(t, "read", permissions[0].Name)
}

func TestPermissionService_ObtainPermissions_Empty(t *testing.T) {
	repo := mocks.NewPermissionRepositoryMock()
	service := NewPermissionService(repo)

	permissions, err := service.ObtainPermissions(context.Background())

	assert.NoError(t, err)
	assert.Len(t, permissions, 0)
}

func TestPermissionService_ObtainPermission_Success(t *testing.T) {
	repo := mocks.NewPermissionRepositoryMock()
	now := time.Now()
	permId := uuid.New()
	perm := &permissionsDomains.Permission{
		Id:          permId,
		Name:        "read",
		Description: "Read permission",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repo.Permissions[permId.String()] = perm

	service := NewPermissionService(repo)
	permToFind := &permissionsDomains.Permission{Id: permId}

	foundPerm, err := service.ObtainPermission(context.Background(), permToFind)

	assert.NoError(t, err)
	assert.NotNil(t, foundPerm)
	assert.Equal(t, "read", foundPerm.Name)
}

func TestPermissionService_ObtainPermission_NotFound(t *testing.T) {
	repo := mocks.NewPermissionRepositoryMock()
	service := NewPermissionService(repo)
	permToFind := &permissionsDomains.Permission{Id: uuid.New()}

	foundPerm, err := service.ObtainPermission(context.Background(), permToFind)

	assert.Error(t, err)
	assert.Nil(t, foundPerm)
}

func TestPermissionService_CreatePermission_Success(t *testing.T) {
	repo := mocks.NewPermissionRepositoryMock()
	service := NewPermissionService(repo)
	perm := &permissionsDomains.Permission{
		Name:        "write",
		Description: "Write permission",
	}

	createdPerm, err := service.CreatePermission(context.Background(), perm)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, createdPerm.Id)
	assert.Equal(t, "write", createdPerm.Name)
}

func TestPermissionService_UpdatePermission_Success(t *testing.T) {
	repo := mocks.NewPermissionRepositoryMock()
	now := time.Now()
	permId := uuid.New()
	perm := &permissionsDomains.Permission{
		Id:          permId,
		Name:        "read",
		Description: "Read permission",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repo.Permissions[permId.String()] = perm

	service := NewPermissionService(repo)
	perm.Description = "Updated description"

	updatedPerm, err := service.UpdatePermission(context.Background(), perm)

	assert.NoError(t, err)
	assert.NotNil(t, updatedPerm)
	assert.Equal(t, "Updated description", updatedPerm.Description)
}

func TestPermissionService_UpdatePermission_NotFound(t *testing.T) {
	repo := mocks.NewPermissionRepositoryMock()
	service := NewPermissionService(repo)
	perm := &permissionsDomains.Permission{
		Id:   uuid.New(),
		Name: "read",
	}

	updatedPerm, err := service.UpdatePermission(context.Background(), perm)

	assert.Error(t, err)
	assert.Nil(t, updatedPerm)
}

func TestPermissionService_DeletePermission_Success(t *testing.T) {
	repo := mocks.NewPermissionRepositoryMock()
	now := time.Now()
	permId := uuid.New()
	perm := &permissionsDomains.Permission{
		Id:          permId,
		Name:        "read",
		Description: "Read permission",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repo.Permissions[permId.String()] = perm

	service := NewPermissionService(repo)

	err := service.DeletePermission(context.Background(), perm)

	assert.NoError(t, err)
	assert.Len(t, repo.Permissions, 0)
}

func TestPermissionService_NewPermissionService(t *testing.T) {
	repo := mocks.NewPermissionRepositoryMock()
	service := NewPermissionService(repo)

	assert.NotNil(t, service)
}
