package db

import (
	"context"
	"embed"
	"finboard/src/core/config"
	"time"

	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed sql/ini.sql
var migrations embed.FS

var Conn *pgxpool.Pool

func InitDB() error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pollConfig, err := pgxpool.ParseConfig(config.Config.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}
	pollConfig.MaxConns = 10
	pollConfig.MinConns = 2
	pollConfig.MaxConnLifetime = 30 * time.Minute
	pollConfig.MaxConnIdleTime = 5 * time.Minute

	pollConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := conn.Exec(context.Background(), "SET TIMEZONE TO '"+config.Config.Timezone+"'")
		if err != nil {
			log.Errorf("Failed to set timezone: %v", err)
			return err
		}
		return err
	}

	Conn, err = pgxpool.NewWithConfig(ctx, pollConfig)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	if err = Conn.Ping(ctx); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	log.Info("Database connection established successfully")

	if err := runMigrations(ctx); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	return nil
}

func runMigrations(ctx context.Context) error {
	sqlFile, err := migrations.Open("sql/ini.sql")
	if err != nil {
		return err
	}
	defer sqlFile.Close()

	sqlBytes, err := migrations.ReadFile("sql/ini.sql")
	if err != nil {
		return err
	}

	_, err = Conn.Exec(ctx, string(sqlBytes))
	if err != nil {
		log.Errorf("Migration warning (may be already applied): %v", err)
		return nil
	}

	log.Info("Migrations applied successfully")
	return nil
}

func GetDB() *pgxpool.Pool {
	return Conn
}
