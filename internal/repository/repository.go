package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func New(log *zap.Logger, db *pgxpool.Pool) Repository {
	return Repository{
		db:  db,
		log: log,
	}
}
