package postgres

import (
	"context"
	"database/sql"
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
		return false, err
	}

	return exists, nil
}

// CreateUser - создает пользователя
func (u *UserReg) CreateUser(ctx context.Context, login, hashedPassword string) (int64, error) {
	var id int64
	query := `INSERT INTO users (login, password_hash, created_at) 
              VALUES ($1, $2, NOW()) 
              RETURNING id`

	err := u.db.QueryRowContext(ctx, query, login, hashedPassword).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
