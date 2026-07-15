package usecase

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/MaximTretjakov/nofelet-web/internal/v1/view"
	"github.com/MaximTretjakov/nofelet-web/middleware/metrics"
	"github.com/MaximTretjakov/nofelet-web/pkg/token"
)

var (
	ErrInvalidCredentials = errors.New("auth invalid login or password")
	ErrTokenGeneration    = errors.New("auth jwt token generation error")
)

// CreateToken - Создание токена доступа
func (uc *UseCase) CreateToken(
	ctx context.Context,
	req view.PostRegistrationRequestData,
) (string, error) {
	if len(req.Login) == 0 && len(req.Password) == 0 {
		return "", fmt.Errorf("auth %w", ErrEmptyCredentials)
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
		return "", ErrTokenGeneration
	}

	metrics.AuthSuccess.Add(ctx, 1)

	return accessesToken, nil
}
