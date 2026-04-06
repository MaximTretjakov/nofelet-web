package web

import (
	"github.com/MaximTretjakov/nofelet-web/internal/dependency"
	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/controller"
	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/usecase"
)

func Register(deps *dependency.Container) {
	c := controller.New(usecase.New())

	r := deps.Web.Routes
	r.GET("api/v1/registration", c.PostRegister)
	r.GET("api/v1/auth", c.PostAuth)
	r.GET("api/v1/logout", c.PostLogout)
}
