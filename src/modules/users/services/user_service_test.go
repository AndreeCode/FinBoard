package services

import (
	"context"
	"finboard/src/mocks"
	usersDomains "finboard/src/modules/users/domains"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserService_ObtainUsers_Success(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	now := time.Now()
	userId := uuid.New()
	roleId := uuid.New()
	user := &usersDomains.User{
		Id:        userId,
		Name:      "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		RoleId:    roleId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.Users[userId.String()] = user

	service := NewUserService(repo)

	users, err := service.ObtainUsers(context.Background())

	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "John", users[0].Name)
	assert.Equal(t, "john@example.com", users[0].Email)
}

func TestUserService_ObtainUsers_Empty(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	service := NewUserService(repo)

	users, err := service.ObtainUsers(context.Background())

	assert.NoError(t, err)
	assert.Len(t, users, 0)
}

func TestUserService_ObtainUser_Success(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	now := time.Now()
	userId := uuid.New()
	roleId := uuid.New()
	user := &usersDomains.User{
		Id:        userId,
		Name:      "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		RoleId:    roleId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.Users[userId.String()] = user

	service := NewUserService(repo)
	userToFind := &usersDomains.User{Id: userId}

	foundUser, err := service.ObtainUser(context.Background(), userToFind)

	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, "John", foundUser.Name)
}

func TestUserService_ObtainUser_NotFound(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	service := NewUserService(repo)
	userToFind := &usersDomains.User{Id: uuid.New()}

	foundUser, err := service.ObtainUser(context.Background(), userToFind)

	assert.Error(t, err)
	assert.Nil(t, foundUser)
}

func TestUserService_CreateUser_Success(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	service := NewUserService(repo)
	roleId := uuid.New()
	user := &usersDomains.User{
		Name:     "John",
		LastName: "Doe",
		Email:    "john@example.com",
		Password: "plainpassword",
		RoleId:   roleId,
	}

	createdUser, err := service.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	assert.NotEqual(t, "", createdUser.Password)
	assert.NotEqual(t, "plainpassword", createdUser.Password)
	assert.Equal(t, "John", createdUser.Name)
}

func TestUserService_UpdateUser_Success(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	now := time.Now()
	userId := uuid.New()
	roleId := uuid.New()
	user := &usersDomains.User{
		Id:        userId,
		Name:      "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		RoleId:    roleId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.Users[userId.String()] = user

	service := NewUserService(repo)
	user.Name = "Jane"
	user.LastName = "Smith"

	updatedUser, err := service.UpdateUser(context.Background(), user)

	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, "Jane", updatedUser.Name)
	assert.Equal(t, "Smith", updatedUser.LastName)
}

func TestUserService_UpdateUser_NotFound(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	service := NewUserService(repo)
	user := &usersDomains.User{
		Id:   uuid.New(),
		Name: "Jane",
	}

	updatedUser, err := service.UpdateUser(context.Background(), user)

	assert.Error(t, err)
	assert.Nil(t, updatedUser)
}

func TestUserService_DeleteUser_Success(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	now := time.Now()
	userId := uuid.New()
	roleId := uuid.New()
	user := &usersDomains.User{
		Id:        userId,
		Name:      "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		RoleId:    roleId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.Users[userId.String()] = user

	service := NewUserService(repo)

	err := service.DeleteUser(context.Background(), user)

	assert.NoError(t, err)
	assert.Len(t, repo.Users, 0)
}

func TestUserService_DeleteUser_NotFound(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	service := NewUserService(repo)
	user := &usersDomains.User{Id: uuid.New()}

	err := service.DeleteUser(context.Background(), user)

	assert.Error(t, err)
}

func TestUserService_NewUserService(t *testing.T) {
	repo := mocks.NewUserRepositoryMock()
	service := NewUserService(repo)

	assert.NotNil(t, service)
}
