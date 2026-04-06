package controller

import (
	"context"

	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/dto"
	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
)

type UseCase interface {
	// UserRegistration - Регистрация нового пользователя
	UserRegistration(ctx context.Context, req view.PostRegistrationRequestData) (dto.UserRegistrationResponse, error)
}
