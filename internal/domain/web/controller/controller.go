package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/usecase"
	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
)

var errUserNotFound = errors.New("user not found")

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) HandleError(err error) (int, view.SimpleErrorBody) {
	switch {
	case errors.Is(err, usecase.ErrEmptyCredentials):
		return http.StatusBadRequest, newError(err)
	case errors.Is(err, usecase.ErrUserExists):
		return http.StatusBadRequest, newError(err)
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound, newError(errUserNotFound)
	}

	return http.StatusInternalServerError, newError(err)
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
