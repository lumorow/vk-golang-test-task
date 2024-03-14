package repository

import "github.com/jmoiron/sqlx"

type FilmPostgres struct {
	db *sqlx.DB
}

func NewFilmPostgres(db *sqlx.DB) Film {
	return &FilmPostgres{
		db: db,
	}
}
