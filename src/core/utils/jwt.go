package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")
)

type JWTConfig struct {
	SecretKey     []byte
	Expiration    time.Duration
	RefreshExpiry time.Duration
}

type Claims struct {
	UserID   uuid.UUID `json:"user_id"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
	RoleID   uuid.UUID `json:"role_id"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

var DefaultJWTConfig = &JWTConfig{
	SecretKey:  []byte("your-secret-key-change-in-production"),
	Expiration: 2 * time.Hour,
}

func GenerateToken(userID uuid.UUID, name, lastName, email string, roleID uuid.UUID, config *JWTConfig) (string, error) {
	if config == nil {
		config = DefaultJWTConfig
	}

	claims := &Claims{
		UserID:   userID,
		Name:     name,
		LastName: lastName,
		Email:    email,
		RoleID:   roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.Expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "finboard",
			Subject:   userID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey)
}

func ValidateAccessToken(tokenString string, config *JWTConfig) (*Claims, error) {
	if config == nil {
		config = DefaultJWTConfig
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return config.SecretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func ValidateRefreshToken(tokenString string, config *JWTConfig) (*RefreshClaims, error) {
	if config == nil {
		config = DefaultJWTConfig
	}

	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return config.SecretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
