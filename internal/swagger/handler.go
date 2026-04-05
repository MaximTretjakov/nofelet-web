package swagger

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/swaggest/swgui/v4cdn"

	"github.com/MaximTretjakov/nofelet-web/config"
)

func Register(r gin.IRouter, swFn func() (*openapi3.T, error), basePath string, config *config.Config) error {
	swDoc, err := swFn()
	if err != nil {
		return err
	}
	var (
		once           sync.Once
		swaggerHandler = v4cdn.NewHandler("App API "+basePath, basePath+"/docs.json", "/")
	)

	r.GET("/documentation/*any", gin.WrapH(swaggerHandler))
	r.GET("/docs.json", func(c *gin.Context) {
		once.Do(func() {
			if config.Debug {
				u, _ := url.Parse(c.Request.Header.Get("Referer"))
				uri := fmt.Sprintf("%s://%s", u.Scheme, u.Hostname())
				if config.AppNamespace == "local" {
					uri = fmt.Sprintf("%s://%s:%s", u.Scheme, u.Hostname(), config.Web.Port)
				}
				swDoc.Servers = append(
					[]*openapi3.Server{
						{
							URL: uri,
						},
					},
					swDoc.Servers[0:]...)
			}
		})
		c.JSON(http.StatusOK, &swDoc)
	})
	return nil
}
