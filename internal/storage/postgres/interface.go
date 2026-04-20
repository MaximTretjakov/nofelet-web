package postgres

import "context"

type User interface {
	// IsLoginUnique - проверяет зарегистрирован ли пользователь
	IsLoginUnique(ctx context.Context, login string) (bool, error)
}
