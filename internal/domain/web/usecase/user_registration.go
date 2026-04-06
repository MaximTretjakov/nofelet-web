package usecase

import (
	"context"

	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/dto"
	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
)

// UserRegistration - Регистрация нового пользователя
func (uc *UseCase) UserRegistration(
	ctx context.Context,
	req view.PostRegistrationRequestData,
) (dto.UserRegistrationResponse, error) {
	// todo реализовать use case или интеграцию
	return dto.UserRegistrationResponse{}, nil
}
