package web

import (
	"github.com/MaximTretjakov/nofelet-web/internal/dependency"
	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/controller"
)

func Register(deps *dependency.Container) {
	c := controller.New(deps.Logger, deps.Cfg)

	r := deps.Web.Routes
	r.GET("api/v1/registration", c.PostRegister)
	r.GET("api/v1/auth", c.PostAuth)
	r.GET("api/v1/deauth", c.PostDeauth)
}
