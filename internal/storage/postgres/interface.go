package postgres

import "context"

type User interface {
	// IsLoginUnique - проверяет зарегистрирован ли пользователь
	IsLoginUnique(ctx context.Context, login string) (bool, error)
	// CreateUser - создает пользователя
	CreateUser(ctx context.Context, login, passwordHash string) (int64, error)
}
