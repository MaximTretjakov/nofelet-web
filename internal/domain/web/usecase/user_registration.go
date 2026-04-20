package usecase

import (
	"context"
	"errors"

	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/dto"
	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
)

var (
	ErrUserExists       = errors.New("user already exists")
	ErrEmptyCredentials = errors.New("empty credentials")
)

// UserRegistration - Регистрация нового пользователя
func (uc *UseCase) UserRegistration(
	ctx context.Context,
	req view.PostRegistrationRequestData,
) (dto.UserRegistrationResponse, error) {
	if len(req.Login) == 0 && len(req.Password) == 0 {
		return dto.UserRegistrationResponse{}, ErrEmptyCredentials
	}

	exist, err := uc.User.IsLoginUnique(ctx, req.Login)
	if err != nil {
		return dto.UserRegistrationResponse{}, err
	}
	if !exist {
		return dto.UserRegistrationResponse{}, ErrUserExists
	}

	// todo CreateUser

	return dto.UserRegistrationResponse{}, nil
}
