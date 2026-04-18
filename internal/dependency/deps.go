package dependency

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/MaximTretjakov/nofelet-web/config"
	"github.com/MaximTretjakov/nofelet-web/internal/dependency/web"
	"github.com/MaximTretjakov/nofelet-web/internal/storage/postgres"
)

// Container основной контейнер зависимостей
type Container struct {
	Web    *web.Container
	Logger *slog.Logger
	Cfg    *config.Config
	DB     *sql.DB
}

func New(Cfg *config.Config, logger *slog.Logger) (*Container, error) {
	WebContainer, err := web.New(Cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("создание сигналинг контейнера: %w", err)
	}

	db, dbErr := postgres.New(Cfg.DB.ConnectionString)
	if dbErr != nil {
		return nil, dbErr
	}

	defer func() {
		if dbErr = db.Close(); err != nil {
			logger.Error("dependency:", slog.Any("db init error:", dbErr))
		}
	}()

	return &Container{
		Web:    WebContainer,
		Logger: logger,
		Cfg:    Cfg,
		DB:     db,
	}, nil
}
