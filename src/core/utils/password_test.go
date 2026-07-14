package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword123"

	hash, err := HashPassword(password, nil)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)
}

func TestHashPasswordWithCustomConfig(t *testing.T) {
	password := "testpassword123"
	config := &PasswordConfig{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := HashPassword(password, config)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestHashPasswordDifferentHashes(t *testing.T) {
	password := "testpassword123"

	hash1, _ := HashPassword(password, nil)
	hash2, _ := HashPassword(password, nil)

	assert.NotEqual(t, hash1, hash2, "Same password should produce different hashes due to random salt")
}

func TestCheckPasswordValid(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password, nil)
	assert.NoError(t, err)

	valid, err := CheckPassword(password, hash)

	assert.NoError(t, err)
	assert.True(t, valid)
}

func TestCheckPasswordInvalid(t *testing.T) {
	password := "testpassword123"
	wrongPassword := "wrongpassword"

	hash, err := HashPassword(password, nil)
	assert.NoError(t, err)

	valid, err := CheckPassword(wrongPassword, hash)

	assert.NoError(t, err)
	assert.False(t, valid)
}

func TestCheckPasswordEmptyPassword(t *testing.T) {
	password := "testpassword123"
	hash, err := HashPassword(password, nil)
	assert.NoError(t, err)

	valid, err := CheckPassword("", hash)

	assert.NoError(t, err)
	assert.False(t, valid)
}

func TestCheckPasswordInvalidHash(t *testing.T) {
	valid, err := CheckPassword("password", "invalid-hash")

	assert.Error(t, err)
	assert.False(t, valid)
	assert.Equal(t, ErrInvalidHash, err)
}

func TestGenerateUUID(t *testing.T) {
	uuid1 := GenerateUUID()
	uuid2 := GenerateUUID()

	assert.NotEmpty(t, uuid1)
	assert.NotEmpty(t, uuid2)
	assert.NotEqual(t, uuid1, uuid2, "UUIDs should be unique")
}

func TestPasswordHashLength(t *testing.T) {
	password := "testpassword"

	hash, err := HashPassword(password, nil)
	assert.NoError(t, err)

	valid, err := CheckPassword(password, hash)
	assert.NoError(t, err)
	assert.True(t, valid)
}

func TestCheckPasswordWithCustomConfig(t *testing.T) {
	password := "testpassword123"
	config := &PasswordConfig{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 4,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := HashPassword(password, config)
	assert.NoError(t, err)

	valid, err := CheckPassword(password, hash)
	assert.NoError(t, err)
	assert.True(t, valid)
}
