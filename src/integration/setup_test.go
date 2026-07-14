package integration

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type IntegrationTestSuite struct {
	DB       *pgxpool.Pool
	Ctx      context.Context
	Teardown func()
}

func SetupPostgresContainer(t *testing.T) (*IntegrationTestSuite, error) {
	ctx := context.Background()

	dbName := fmt.Sprintf("finboard_test_%d", time.Now().UnixNano())

	container, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to start postgres container: %w", err)
	}

	connStr, err := container.ConnectionString(ctx)
	if err != nil {
		container.Terminate(ctx)
		return nil, fmt.Errorf("failed to get connection string: %w", err)
	}

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		container.Terminate(ctx)
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		container.Terminate(ctx)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Printf("Connected to PostgreSQL container: %s", shortID(container))

	return &IntegrationTestSuite{
		DB:  pool,
		Ctx: ctx,
		Teardown: func() {
			pool.Close()
			if err := container.Terminate(ctx); err != nil {
				log.Printf("Error terminating container: %v", err)
			}
		},
	}, nil
}

func shortID(c testcontainers.Container) string {
	id := c.GetContainerID()
	if len(id) > 8 {
		return id[:8]
	}
	return id
}

func (s *IntegrationTestSuite) RunMigrations(t *testing.T) {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			lastname VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			role_id UUID,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			created_by UUID,
			last_login TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS roles (
			id UUID PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			created_by UUID
		)`,
		`CREATE TABLE IF NOT EXISTS permissions (
			id UUID PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			created_by UUID
		)`,
		`CREATE TABLE IF NOT EXISTS categories (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			parent_id UUID,
			user_id UUID,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			created_by UUID
		)`,
		`CREATE TABLE IF NOT EXISTS transactions (
			id UUID PRIMARY KEY,
			user_id UUID NOT NULL,
			category_id UUID,
			amount DECIMAL(15,2) NOT NULL,
			type VARCHAR(50) NOT NULL,
			transaction_date TIMESTAMP NOT NULL,
			received_date TIMESTAMP,
			due_date TIMESTAMP,
			canceled BOOLEAN DEFAULT FALSE,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			created_by UUID
		)`,
		`CREATE TABLE IF NOT EXISTS credits (
			id UUID PRIMARY KEY,
			user_id UUID NOT NULL,
			person_name VARCHAR(255) NOT NULL,
			amount DECIMAL(15,2) NOT NULL,
			interest_rate DECIMAL(5,2),
			is_creditor BOOLEAN DEFAULT TRUE,
			is_secure BOOLEAN DEFAULT FALSE,
			due_date TIMESTAMP,
			status VARCHAR(50) DEFAULT 'active',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			created_by UUID
		)`,
		`CREATE TABLE IF NOT EXISTS investments (
			id UUID PRIMARY KEY,
			transaction_id UUID NOT NULL,
			expected_gain DECIMAL(15,2),
			risk_level VARCHAR(50),
			status VARCHAR(50) DEFAULT 'active',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			created_by UUID
		)`,
		`CREATE TABLE IF NOT EXISTS role_permissions (
			id UUID PRIMARY KEY,
			role_id UUID NOT NULL,
			permission_id UUID NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP,
			created_by UUID
		)`,
	}

	for _, migration := range migrations {
		if _, err := s.DB.Exec(s.Ctx, migration); err != nil {
			t.Fatalf("Migration failed: %v", err)
		}
	}
}
