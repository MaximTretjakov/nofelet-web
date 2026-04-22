package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const dbType = "pgx"

func New(connection string) (*sql.DB, error) {
	db, err := sql.Open(dbType, connection)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	// Макс. кол-во активных соединений
	db.SetMaxOpenConns(5)
	// Макс. кол-во соединений в ожидании
	db.SetMaxIdleConns(2)
	// Время жизни соединения
	db.SetConnMaxLifetime(5 * time.Minute)
	// Время жизни простаивающего соединения
	db.SetConnMaxIdleTime(2 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}
