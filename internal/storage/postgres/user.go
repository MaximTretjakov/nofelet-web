package postgres

import (
	"context"
	"database/sql"

	"github.com/MaximTretjakov/nofelet-web/internal/domain/web/dto"
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
func (u *UserReg) CreateUser(ctx context.Context, login, hashedPassword, name string) (int64, error) {
	var id int64
	query := `INSERT INTO users (login, password, name, created_at) 
              VALUES ($1, $2, $3, NOW()) 
              RETURNING id`

	err := u.db.QueryRowContext(ctx, query, login, hashedPassword, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetUserCredentials - извлекает креды пользователя
func (u *UserReg) GetUserCredentials(ctx context.Context, login string) (dto.UserCredentials, error) {
	var user dto.UserCredentials
	query := `SELECT login, password FROM users WHERE login = $1`

	err := u.db.QueryRowContext(ctx, query, login).Scan(&user.Login, &user.Password)
	if err != nil {
		return dto.UserCredentials{}, err
	}

	return user, nil
}
