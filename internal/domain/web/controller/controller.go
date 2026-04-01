package controller

import (
	"log/slog"

	"github.com/MaximTretjakov/nofelet-web/config"
)

type Controller struct {
	Logger *slog.Logger
	Config *config.Config
}

func New(logger *slog.Logger, cfg *config.Config) *Controller {
	return &Controller{
		Logger: logger,
		Config: cfg,
	}
}
