package repository

import (
	"filmlib/server/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type Actor interface {
	CreateActor(actor entity.Actor) (int, error)
	DeleteActorById(actorId int) error
	UpdateActorById(actorId int, actor entity.UpdateActorInput) error
	GetActor(actorId int) (entity.Actor, error)
	GetActorsIdByFilmId(filmId int) ([]int, error)
}

type Film interface {
	CreateFilm(film entity.Film) (int, error)
	DeleteFilmById(filmId int) error
	GetFilmsByActorId(actorId int) ([]entity.Film, error)
	GetFilmsWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error)
	GetFilmsWithSort(sortType string, filmsId []int) ([]entity.Film, error)
	UpdateFilmById(filmId int, deleteIds []int, addIds []int, film entity.UpdateFilmInput) error
}

type Repository struct {
	Authorization
	Actor
	Film
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Actor:         NewActorPostgres(db),
		Film:          NewFilmPostgres(db),
	}
}
