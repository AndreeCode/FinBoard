package services

import (
	"context"
	"errors"
	"finboard/src/core/utils"
	authDomains "finboard/src/modules/auth/domains"
	userDomains "finboard/src/modules/users/domains"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrTokenGeneration    = errors.New("error generating token")
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*userDomains.User, error)
	UpdateLastLogin(ctx context.Context, user *userDomains.User) error
}

type AuthService struct {
	userRepo UserRepositoryInterface
}

func NewAuthService(userRepo UserRepositoryInterface) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(ctx context.Context, req *authDomains.LoginRequest) (*authDomains.AuthResponse, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	valid, err := utils.CheckPassword(req.Password, user.Password)
	if err != nil || !valid {
		return nil, ErrInvalidCredentials
	}

	err = s.userRepo.UpdateLastLogin(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.Id, user.Name, user.LastName, user.Email, user.RoleId, nil)
	if err != nil {
		return nil, ErrTokenGeneration
	}

	return &authDomains.AuthResponse{
		Token: token,
		User: map[string]interface{}{
			"id":        user.Id,
			"name":      user.Name,
			"last_name": user.LastName,
			"email":     user.Email,
			"role_id":   user.RoleId,
		},
	}, nil
}

func (s *AuthService) ValidateToken(accessToken string) (*utils.Claims, error) {
	claims, err := utils.ValidateAccessToken(accessToken, nil)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (s *AuthService) HashPassword(password string) (string, error) {
	return utils.HashPassword(password, nil)
}
