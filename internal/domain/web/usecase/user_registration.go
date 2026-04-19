package usecase

import (
	"context"
	"errors"

	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/dto"
	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
)

var ErrEmptyCredentials = errors.New("empty credentials")

// UserRegistration - Регистрация нового пользователя
func (uc *UseCase) UserRegistration(
	ctx context.Context,
	req view.PostRegistrationRequestData,
) (dto.UserRegistrationResponse, error) {
	if len(req.Login) == 0 && len(req.Password) == 0 {
		return dto.UserRegistrationResponse{}, ErrEmptyCredentials
	}

	return dto.UserRegistrationResponse{}, nil
}
