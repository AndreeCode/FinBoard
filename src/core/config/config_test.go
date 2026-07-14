package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "secret")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("PORT", "5000")
	os.Setenv("JWT_SECRET", "test-jwt-secret")
	os.Setenv("TIMEZONE", "America/Lima")

	LoadConfig()

	assert.Equal(t, "localhost", Config.DBHost)
	assert.Equal(t, "5432", Config.DBPort)
	assert.Equal(t, "postgres", Config.DBUser)
	assert.Equal(t, "secret", Config.DBPassword)
	assert.Equal(t, "testdb", Config.DBName)
	assert.Equal(t, "5000", Config.Port)
	assert.Equal(t, "test-jwt-secret", Config.JWTSecret)
	assert.Equal(t, "America/Lima", Config.Timezone)
}

func TestLoadConfigDatabaseURL(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "secret")
	os.Setenv("DB_NAME", "testdb")

	LoadConfig()

	expectedURL := " user = postgres password= secret host= localhost port= 5432 dbname= testdb sslmode=disable"
	assert.Equal(t, expectedURL, Config.DatabaseURL)
}

func TestLoadConfigEmptyValues(t *testing.T) {
	os.Setenv("DB_HOST", "")
	os.Setenv("DB_PORT", "")
	os.Setenv("DB_USER", "")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("DB_NAME", "")

	LoadConfig()

	assert.Equal(t, "", Config.DBHost)
	assert.Equal(t, "", Config.DBPort)
	assert.Equal(t, "", Config.DBUser)
	assert.Equal(t, "", Config.DBPassword)
	assert.Equal(t, "", Config.DBName)
}

func TestAppConfigStructure(t *testing.T) {
	cfg := AppConfig{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "password",
		DBName:     "finboard",
		Timezone:   "America/Lima",
		Port:       "5000",
		JWTSecret:  "secret",
	}

	assert.Equal(t, "localhost", cfg.DBHost)
	assert.Equal(t, "5432", cfg.DBPort)
	assert.Equal(t, "postgres", cfg.DBUser)
	assert.Equal(t, "password", cfg.DBPassword)
	assert.Equal(t, "finboard", cfg.DBName)
	assert.Equal(t, "America/Lima", cfg.Timezone)
	assert.Equal(t, "5000", cfg.Port)
	assert.Equal(t, "secret", cfg.JWTSecret)
}
