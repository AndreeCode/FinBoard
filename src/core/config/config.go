package config

import (
	"os"

	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DatabaseURL string
	Timezone    string

	Port      string
	JWTSecret string
}

var Config AppConfig

func LoadConfig() {
	_ = godotenv.Load()

	Config = AppConfig{
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Timezone:    os.Getenv("TIMEZONE"),

		Port:      os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	Config.DatabaseURL =
		" user = " + Config.DBUser +
			" password= " + Config.DBPassword +
			" host= " + Config.DBHost +
			" port= " + Config.DBPort +
			" dbname= " + Config.DBName +
			" sslmode=disable"

	log.Info("Configuration loaded successfully")
}
