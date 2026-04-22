package usecase

import (
	"database/sql"
	"log/slog"

	"github.com/MaximTretjakov/nofelet-web/internal/storage/postgres"
)

type UseCase struct {
	DB   *sql.DB
	Log  *slog.Logger
	User postgres.User
}

func New(db *sql.DB, log *slog.Logger, user *postgres.UserReg) *UseCase {
	return &UseCase{
		DB:   db,
		Log:  log,
		User: user,
	}
}
