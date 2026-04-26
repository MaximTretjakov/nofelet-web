package postgres

import (
	"context"

	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/dto"
)

type User interface {
	// IsLoginUnique - проверяет зарегистрирован ли пользователь
	IsLoginUnique(ctx context.Context, login string) (bool, error)
	// CreateUser - создает пользователя
	CreateUser(ctx context.Context, login, passwordHash, name string) (int64, error)
	// GetUserCredentials - извлекает креды пользователя
	GetUserCredentials(ctx context.Context, login string) (dto.UserCredentials, error)
}
