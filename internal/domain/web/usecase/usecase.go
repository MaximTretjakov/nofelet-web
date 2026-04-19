package usecase

import (
	"database/sql"
	"log/slog"
)

type UseCase struct {
	DB  *sql.DB
	Log *slog.Logger
}

func New(db *sql.DB, log *slog.Logger) *UseCase {
	return &UseCase{
		DB:  db,
		Log: log,
	}
}
