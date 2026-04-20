package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type UserReg struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *UserReg {
	return &UserReg{
		db: db,
	}
}

// IsLoginUnique - проверяет зарегистрирован ли пользователь
func (u *UserReg) IsLoginUnique(ctx context.Context, login string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)`

	err := u.db.QueryRowContext(ctx, query, login).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check login uniqueness error: %w", err)
	}

	return !exists, nil
}
