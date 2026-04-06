package controller

import (
	"errors"
	"net/http"

	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
	clientErrors "github.com/MaximTretjakov/nofelet-web/services/client/errors"
)

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) HandleError(err error) (int, view.SimpleErrorBody) {
	var se clientErrors.ServiceError
	if errors.As(err, &se) {
		return se.StatusCode, newError(se)
	}
	return http.StatusInternalServerError, view.SimpleErrorBody{}
}

func newError(err error) view.SimpleErrorBody {
	return view.SimpleErrorBody{
		Data: struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		},
	}
}
