package services

import (
	"context"
	"errors"
	"finboard/src/core/utils"
	authDomains "finboard/src/modules/auth/domains"
	userDomains "finboard/src/modules/users/domains"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct {
	users          map[string]*userDomains.User
	updateErr      error
	updateLastErr   error
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{
		users: make(map[string]*userDomains.User),
	}
}

func (m *mockUserRepo) GetUserByEmail(ctx context.Context, email string) (*userDomains.User, error) {
	for _, u := range m.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, ErrUserNotFound
}

func (m *mockUserRepo) UpdateLastLogin(ctx context.Context, user *userDomains.User) error {
	if m.updateLastErr != nil {
		return m.updateLastErr
	}
	now := time.Now()
	user.LastLogin = &now
	return nil
}

func TestAuthService_Login_Success(t *testing.T) {
	repo := newMockUserRepo()
	now := time.Now()
	userId := uuid.New()
	roleId := uuid.New()
	password := "password123"

	hashedPassword, _ := utils.HashPassword(password, nil)

	user := &userDomains.User{
		Id:        userId,
		Name:      "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  hashedPassword,
		RoleId:    roleId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.users[userId.String()] = user

	authService := &AuthService{userRepo: repo}

	loginReq := &authDomains.LoginRequest{
		Email:    "john@example.com",
		Password: password,
	}

	resp, err := authService.Login(context.Background(), loginReq)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Token)
	userMap := resp.User.(map[string]interface{})
	assert.Equal(t, "John", userMap["name"])
	assert.Equal(t, "john@example.com", userMap["email"])
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	repo := newMockUserRepo()
	authService := &AuthService{userRepo: repo}

	loginReq := &authDomains.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	resp, err := authService.Login(context.Background(), loginReq)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, ErrUserNotFound, err)
}

func TestAuthService_Login_InvalidPassword(t *testing.T) {
	repo := newMockUserRepo()
	now := time.Now()
	userId := uuid.New()
	roleId := uuid.New()

	user := &userDomains.User{
		Id:        userId,
		Name:      "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		RoleId:    roleId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.users[userId.String()] = user

	authService := &AuthService{userRepo: repo}

	loginReq := &authDomains.LoginRequest{
		Email:    "john@example.com",
		Password: "wrongpassword",
	}

	resp, err := authService.Login(context.Background(), loginReq)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, ErrInvalidCredentials, err)
}

func TestAuthService_ValidateToken_Success(t *testing.T) {
	repo := newMockUserRepo()
	authService := &AuthService{userRepo: repo}

	userId := uuid.New()
	roleId := uuid.New()

	token, _ := utils.GenerateToken(userId, "John", "Doe", "john@example.com", roleId, nil)

	claims, err := authService.ValidateToken(token)

	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, userId, claims.UserID)
	assert.Equal(t, "John", claims.Name)
	assert.Equal(t, "Doe", claims.LastName)
	assert.Equal(t, "john@example.com", claims.Email)
}

func TestAuthService_ValidateToken_Invalid(t *testing.T) {
	repo := newMockUserRepo()
	authService := &AuthService{userRepo: repo}

	claims, err := authService.ValidateToken("invalid-token")

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestAuthService_HashPassword(t *testing.T) {
	repo := newMockUserRepo()
	authService := &AuthService{userRepo: repo}

	password := "testpassword123"
	hashedPassword, err := authService.HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
	assert.NotEqual(t, password, hashedPassword)

	valid, err := utils.CheckPassword(password, hashedPassword)
	assert.NoError(t, err)
	assert.True(t, valid)
}

func TestAuthServiceErrors(t *testing.T) {
	assert.NotNil(t, ErrInvalidCredentials)
	assert.NotNil(t, ErrUserNotFound)
	assert.NotNil(t, ErrTokenGeneration)

	assert.Equal(t, "invalid credentials", ErrInvalidCredentials.Error())
	assert.Equal(t, "user not found", ErrUserNotFound.Error())
	assert.Equal(t, "error generating token", ErrTokenGeneration.Error())
}

func TestAuthService_NewAuthService(t *testing.T) {
	repo := newMockUserRepo()
	authService := NewAuthService(repo)

	assert.NotNil(t, authService)
	assert.Equal(t, repo, authService.userRepo)
}

func TestAuthService_Login_UpdateLastLoginError(t *testing.T) {
	repo := newMockUserRepo()
	repo.updateLastErr = errors.New("database error")
	now := time.Now()
	userId := uuid.New()
	roleId := uuid.New()
	password := "password123"

	hashedPassword, _ := utils.HashPassword(password, nil)

	user := &userDomains.User{
		Id:        userId,
		Name:      "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  hashedPassword,
		RoleId:    roleId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	repo.users[userId.String()] = user

	authService := &AuthService{userRepo: repo}

	loginReq := &authDomains.LoginRequest{
		Email:    "john@example.com",
		Password: password,
	}

	resp, err := authService.Login(context.Background(), loginReq)

	assert.Error(t, err)
	assert.Nil(t, resp)
}
