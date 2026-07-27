package usecase

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"golang.org/x/crypto/bcrypt"

	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/dto"
	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
	"github.com/MaximTretjakov/nofelet-web/middleware/metrics"
)

var (
	ErrUserExists       = errors.New("user already exists")
	ErrEmptyCredentials = errors.New("empty credentials")
	ErrHashingFailed    = errors.New("failed to hash password")
)

// UserRegistration - Регистрация нового пользователя
func (uc *UseCase) UserRegistration(
	ctx context.Context,
	req view.PostRegistrationRequestData,
) (dto.UserRegistrationResponse, error) {
	if len(req.Login) == 0 && len(req.Password) == 0 && len(req.Name) == 0 {
		return dto.UserRegistrationResponse{}, ErrEmptyCredentials
	}

	exist, err := uc.User.IsLoginUnique(ctx, req.Login)
	if err != nil {
		return dto.UserRegistrationResponse{}, err
	}
	if exist {
		return dto.UserRegistrationResponse{}, ErrUserExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserRegistrationResponse{}, ErrHashingFailed
	}
	hashedPassword := string(hash)

	_, err = uc.User.CreateUser(ctx, req.Login, hashedPassword, req.Name)
	if err != nil {
		metrics.UserCreationFail.Add(
			ctx,
			1,
			metric.WithAttributes(attribute.String("reason", err.Error())),
		)
		return dto.UserRegistrationResponse{}, err
	}

	metrics.UserCreationSuccess.Add(ctx, 1)
	metrics.RegisterSuccess.Add(ctx, 1)

	return dto.UserRegistrationResponse{}, nil
}
