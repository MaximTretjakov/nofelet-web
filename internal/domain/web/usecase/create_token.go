package usecase

import (
	"context"
	"errors"
	"log/slog"

	"golang.org/x/crypto/bcrypt"

	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
	"github.com/MaximTretjakov/nofelet-web/pkg/token"
)

var ErrInvalidCredentials = errors.New("invalid login or password")

// CreateToken - Создание токена доступа
func (uc *UseCase) CreateToken(
	ctx context.Context,
	req view.PostRegistrationRequestData,
) (string, error) {
	if len(req.Login) == 0 && len(req.Password) == 0 {
		return "", ErrEmptyCredentials
	}

	cred, cErr := uc.User.GetUserCredentials(ctx, req.Login)
	if cErr != nil {
		return "", cErr
	}

	hErr := bcrypt.CompareHashAndPassword([]byte(cred.Password), []byte(req.Password))
	if hErr != nil {
		return "", ErrInvalidCredentials
	}

	model := token.NewJWTToken(uc.Cfg)
	accessesToken, tErr := model.CreateToken(cred)
	if tErr != nil {
		uc.Log.Error("createToken:", slog.Any("ошибка генерации токена:", tErr))
		return "", ErrInvalidCredentials
	}

	return accessesToken, nil
}
