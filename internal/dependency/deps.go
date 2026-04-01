package dependency

import (
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
}

func New(Cfg *config.Config, logger *slog.Logger) (*Container, error) {
	WebContainer, err := web.New(Cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("создание сигналинг контейнера: %w", err)
	}

	return &Container{
		Web:    WebContainer,
		Logger: logger,
		Cfg:    Cfg,
	}, nil
}
