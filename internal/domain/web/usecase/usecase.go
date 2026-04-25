package usecase

import (
	"database/sql"
	"log/slog"

	"github.com/MaximTretjakov/nofelet-web/config"
	"github.com/MaximTretjakov/nofelet-web/internal/storage/postgres"
)

type UseCase struct {
	DB   *sql.DB
	Cfg  *config.Config
	Log  *slog.Logger
	User postgres.User
}

func New(db *sql.DB, cfg *config.Config, log *slog.Logger, user *postgres.UserReg) *UseCase {
	return &UseCase{
		DB:   db,
		Cfg:  cfg,
		Log:  log,
		User: user,
	}
}
