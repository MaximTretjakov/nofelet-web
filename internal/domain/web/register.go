package web

import (
	"github.com/MaximTretjakov/nofelet-web/internal/dependency"
	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/controller"
	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/usecase"
)

func Register(deps *dependency.Container) {
	c := controller.New(usecase.New())

	r := deps.Web.Routes.Group("/nofelet-web/api/v1")
	r.POST("/registration", c.PostRegister)
	r.POST("/auth", c.PostAuth)
	r.PUT("/logout", c.PostLogout)
}
