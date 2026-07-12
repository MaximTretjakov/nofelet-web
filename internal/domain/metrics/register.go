package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/MaximTretjakov/nofelet-web/internal/dependency"
)

func Register(deps *dependency.Container) {
	r := deps.Web.Routes.Group("/nofelet-web/api/v1")
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
