package web

import (
	"github.com/MaximTretjakov/nofelet-web/internal/dependency"
	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/controller"
	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/usecase"
	"github.com/MaximTretjakov/nofelet-web/internal/storage/postgres"
)

func Register(deps *dependency.Container) {
	c := controller.New(makeUC(deps))

	r := deps.Web.Routes.Group("/nofelet-web/api/v1")
	r.POST("/registration", c.PostRegister)
	r.POST("/auth", c.PostAuth)
	r.POST("/logout", c.PostLogout)
}

func makeUC(deps *dependency.Container) *usecase.UseCase {
	return usecase.New(
		deps.DB,
		deps.Cfg,
		deps.Logger,
		postgres.NewUser(deps.DB),
	)
}
