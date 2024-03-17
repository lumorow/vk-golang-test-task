package repository

import (
	"database/sql"
)

type DBTX interface {
	QueryRow(query string, args ...any) *sql.Row
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	Begin() (*sql.Tx, error)
	Select(dest interface{}, query string, args ...interface{}) error
}

type Repository struct {
	db DBTX
}

func NewRepository(db DBTX) *Repository {
	return &Repository{
		db,
	}
}
