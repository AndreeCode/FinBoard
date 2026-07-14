package repository

import (
	"finboard/src/core/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateRepository struct {
	DB *pgxpool.Pool
}

func NewCreateRepository() *CreateRepository {
	return &CreateRepository{
		DB: db.GetDB(),
	}
}
