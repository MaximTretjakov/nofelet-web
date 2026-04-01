package web

import (
	"github.com/MaximTretjakov/nofelet-web/internal/dependency"
	web "github.com/MaximTretjakov/nofelet-web/internal/domain/web"
)

func New(deps *dependency.Container) error {
	web.Register(deps)

	return nil
}
