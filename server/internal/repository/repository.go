package repository

import (
	"database/sql"
)

//go:generate mockgen -destination=mocks/repository.go -package=mock -source=repository.go
//go:generate touch mocks/.coverignore

type DBTX interface {
	QueryRow(query string, args ...any) *sql.Row
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	Begin() (*sql.Tx, error)
	Select(dest interface{}, query string, args ...interface{}) error
}

type AuthorizationRepository struct {
	db DBTX
}

type ActorRepository struct {
	db DBTX
}

type FilmRepository struct {
	db DBTX
}

func NewAuthorizationRepository(db DBTX) *AuthorizationRepository {
	return &AuthorizationRepository{
		db,
	}
}

func NewActorRepository(db DBTX) *ActorRepository {
	return &ActorRepository{
		db,
	}
}

func NewFilmRepository(db DBTX) *FilmRepository {
	return &FilmRepository{
		db,
	}
}
