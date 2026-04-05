package web

import (
	"github.com/MaximTretjakov/nofelet-web/internal/dependency"
	"github.com/MaximTretjakov/nofelet-web/internal/domain/web"
	"github.com/MaximTretjakov/nofelet-web/internal/swagger"
	"github.com/MaximTretjakov/nofelet-web/internal/v1"
)

func New(deps *dependency.Container) error {
	if err := initializeSwagger(deps); err != nil {
		return err
	}

	web.Register(deps)

	return nil
}

// initializeSwagger - инициализация документации
func initializeSwagger(deps *dependency.Container) (err error) {
	err = swagger.Register(deps.Web.Routes.Group("/api/v1"), v1.GetSwagger, "/api/v1", deps.Cfg)
	if err != nil {
		return err
	}
	return nil
}
