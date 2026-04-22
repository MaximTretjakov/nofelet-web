package dependency

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/MaximTretjakov/nofelet-web/config"
	"github.com/MaximTretjakov/nofelet-web/internal/dependency/web"
)

// Container основной контейнер зависимостей
type Container struct {
	Web    *web.Container
	Logger *slog.Logger
	Cfg    *config.Config
	DB     *sql.DB
}

func New(Cfg *config.Config, logger *slog.Logger, db *sql.DB) (*Container, error) {
	WebContainer, err := web.New(Cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("создание сигналинг контейнера: %w", err)
	}

	return &Container{
		Web:    WebContainer,
		Logger: logger,
		Cfg:    Cfg,
		DB:     db,
	}, nil
}
