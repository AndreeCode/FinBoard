package utils

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	userID := uuid.New()
	roleID := uuid.New()

	token, err := GenerateToken(userID, "John", "Doe", "john@example.com", roleID, nil)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGenerateTokenWithCustomConfig(t *testing.T) {
	userID := uuid.New()
	roleID := uuid.New()
	config := &JWTConfig{
		SecretKey:  []byte("test-secret-key"),
		Expiration: 1 * time.Hour,
	}

	token, err := GenerateToken(userID, "John", "Doe", "john@example.com", roleID, config)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateAccessToken(t *testing.T) {
	userID := uuid.New()
	roleID := uuid.New()

	token, err := GenerateToken(userID, "John", "Doe", "john@example.com", roleID, nil)
	assert.NoError(t, err)

	claims, err := ValidateAccessToken(token, nil)

	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
	assert.Equal(t, "John", claims.Name)
	assert.Equal(t, "Doe", claims.LastName)
	assert.Equal(t, "john@example.com", claims.Email)
	assert.Equal(t, roleID, claims.RoleID)
}

func TestValidateAccessTokenWithCustomConfig(t *testing.T) {
	userID := uuid.New()
	roleID := uuid.New()
	config := &JWTConfig{
		SecretKey:  []byte("test-secret-key"),
		Expiration: 1 * time.Hour,
	}

	token, err := GenerateToken(userID, "John", "Doe", "john@example.com", roleID, config)
	assert.NoError(t, err)

	claims, err := ValidateAccessToken(token, config)

	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
}

func TestValidateAccessTokenInvalidToken(t *testing.T) {
	claims, err := ValidateAccessToken("invalid-token", nil)

	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Equal(t, ErrInvalidToken, err)
}

func TestValidateAccessTokenEmptyToken(t *testing.T) {
	claims, err := ValidateAccessToken("", nil)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateAccessTokenWrongSecret(t *testing.T) {
	userID := uuid.New()
	roleID := uuid.New()
	config := &JWTConfig{
		SecretKey:  []byte("correct-secret-key-1234567890"),
		Expiration: 1 * time.Hour,
	}

	token, err := GenerateToken(userID, "John", "Doe", "john@example.com", roleID, config)
	assert.NoError(t, err)

	wrongConfig := &JWTConfig{
		SecretKey:  []byte("wrong-secret-key-1234567890"),
		Expiration: 1 * time.Hour,
	}

	claims, err := ValidateAccessToken(token, wrongConfig)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateRefreshToken(t *testing.T) {
	userID := uuid.New()
	config := &JWTConfig{
		SecretKey:      []byte("test-secret-key-1234567890123456"),
		Expiration:     1 * time.Hour,
		RefreshExpiry:  7 * 24 * time.Hour,
	}

	token, err := GenerateToken(userID, "John", "Doe", "john@example.com", uuid.New(), config)
	assert.NoError(t, err)

	claims, err := ValidateRefreshToken(token, config)

	assert.NoError(t, err)
	assert.Equal(t, userID, claims.UserID)
}

func TestValidateRefreshTokenInvalid(t *testing.T) {
	claims, err := ValidateRefreshToken("invalid-refresh-token", nil)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestClaims_Fields(t *testing.T) {
	userID := uuid.New()
	roleID := uuid.New()
	config := &JWTConfig{
		SecretKey:  []byte("test-secret-key-1234567890123456"),
		Expiration: 2 * time.Hour,
	}

	token, err := GenerateToken(userID, "John", "Doe", "john@example.com", roleID, config)
	assert.NoError(t, err)

	claims, err := ValidateAccessToken(token, config)
	assert.NoError(t, err)

	assert.Equal(t, "finboard", claims.Issuer)
	assert.Equal(t, userID.String(), claims.Subject)
	assert.NotNil(t, claims.ExpiresAt)
	assert.NotNil(t, claims.IssuedAt)
	assert.NotNil(t, claims.NotBefore)
}
