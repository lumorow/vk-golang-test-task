package repository

import "github.com/jmoiron/sqlx"

type ActorPostgres struct {
	db *sqlx.DB
}

func NewActorPostgres(db *sqlx.DB) Actor {
	return &ActorPostgres{
		db: db,
	}
}
