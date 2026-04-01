package web

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/MaximTretjakov/nofelet-web/config"
	"github.com/MaximTretjakov/nofelet-web/middleware"
)

type Container struct {
	Routes *gin.Engine
	Logger *slog.Logger
	Cfg    *config.Config
}

func New(cfg *config.Config, logger *slog.Logger) (*Container, error) {
	routes, err := newRoutes()
	if err != nil {
		return nil, fmt.Errorf("инициализация роутера: %w", err)
	}

	return &Container{
		Routes: routes,
		Logger: logger,
		Cfg:    cfg,
	}, nil
}

func newRoutes() (*gin.Engine, error) {
	router := gin.New()
	router.ContextWithFallback = true
	router.HandleMethodNotAllowed = true
	router.Use(
		gin.Recovery(),
		middleware.DurationLoggerMiddleware(),
	)
	return router, nil
}
