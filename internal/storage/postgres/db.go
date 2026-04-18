package postgres

import (
	"database/sql"
	"time"
)

const dbType = "postgres"

func New(connection string) (*sql.DB, error) {
	db, err := sql.Open(dbType, connection)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
