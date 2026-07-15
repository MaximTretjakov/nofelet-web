package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/usecase"
	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
	"github.com/MaximTretjakov/nofelet-web/middleware/metrics"
)

var (
	errUserNotFound         = errors.New("user not found")
	errAuthEmptyCredentials = errors.New("auth empty credentials")
)

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) HandleError(ctx *gin.Context, err error) (int, view.SimpleErrorBody) {
	switch {
	case errors.Is(err, usecase.ErrEmptyCredentials):
		metrics.RegisterFail.Add(
			ctx,
			1,
			metric.WithAttributes(
				attribute.String("reason", err.Error()),
			))
		return http.StatusBadRequest, newError(err)
	case errors.Is(err, usecase.ErrUserExists):
		metrics.RegisterFail.Add(
			ctx,
			1,
			metric.WithAttributes(
				attribute.String("reason", err.Error()),
			))
		return http.StatusBadRequest, newError(err)
	case errors.Is(err, sql.ErrNoRows):
		metrics.RegisterFail.Add(
			ctx,
			1,
			metric.WithAttributes(
				attribute.String("reason", errUserNotFound.Error()),
			))
		return http.StatusNotFound, newError(errUserNotFound)
	case errors.Is(err, errAuthEmptyCredentials):
		metrics.AuthFail.Add(
			ctx,
			1,
			metric.WithAttributes(
				attribute.String("reason", errAuthEmptyCredentials.Error()),
			))
		return http.StatusNotFound, newError(errAuthEmptyCredentials)
	case errors.Is(err, usecase.ErrInvalidCredentials):
		metrics.AuthFail.Add(
			ctx,
			1,
			metric.WithAttributes(
				attribute.String("reason", usecase.ErrInvalidCredentials.Error()),
			))
		return http.StatusNotFound, newError(usecase.ErrInvalidCredentials)
	case errors.Is(err, usecase.ErrTokenGeneration):
		metrics.AuthFail.Add(
			ctx,
			1,
			metric.WithAttributes(
				attribute.String("reason", usecase.ErrTokenGeneration.Error()),
			))
		return http.StatusNotFound, newError(usecase.ErrTokenGeneration)
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
