package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("invalid hash")
	ErrIncompatibleVersion = errors.New("incompatible version")
)

type PasswordConfig struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var DefaultPasswordConfig = &PasswordConfig{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 4,
	SaltLength:  16,
	KeyLength:   32,
}

func HashPassword(password string, config *PasswordConfig) (string, error) {
	if config == nil {
		config = DefaultPasswordConfig
	}

	salt := make([]byte, config.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		config.Iterations,
		config.Memory,
		config.Parallelism,
		config.KeyLength,
	)

	raw := make([]byte, 0, 1+len(salt)+len(hash))
	raw = append(raw, argon2.Version)
	raw = append(raw, salt...)
	raw = append(raw, hash...)

	return base64.StdEncoding.EncodeToString(raw), nil
}

func CheckPassword(password, encodedHash string) (bool, error) {
	cfg := DefaultPasswordConfig
	if cfg == nil {
		cfg = DefaultPasswordConfig
	}

	b, err := base64.StdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false, ErrInvalidHash
	}

	if len(b) < 1+int(cfg.SaltLength)+1 {
		return false, ErrInvalidHash
	}

	version := int(b[0])
	if version != argon2.Version {
		return false, ErrIncompatibleVersion
	}

	salt := b[1 : 1+cfg.SaltLength]
	storedHash := b[1+cfg.SaltLength:]

	if len(salt) != int(cfg.SaltLength) || len(storedHash) != int(cfg.KeyLength) {
		return false, ErrInvalidHash
	}

	computedHash := argon2.IDKey(
		[]byte(password),
		salt,
		cfg.Iterations,
		cfg.Memory,
		cfg.Parallelism,
		cfg.KeyLength,
	)

	return string(storedHash) == string(computedHash), nil
}

func GenerateUUID() string {
	return uuid.New().String()
}
