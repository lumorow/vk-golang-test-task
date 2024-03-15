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
	DeleteActorById(id int, actor entity.Actor) error
	UpdateActorById(id int, actor entity.Actor) error
	GetActorsWithFilms(id []int) ([]entity.ActorFilms, error)
}

type Film interface {
	CreateFilm(actor entity.Film) (int, error)
	DeleteFilmById(id int, actor entity.Film) error
	UpdateFilmById(id int, actor entity.Film) error
	GetFilmsByActorId(actorId int) ([]entity.Film, error)
	GetFilmsWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error)
	GetFilmsWithSort(sortMode string) ([]entity.Film, error)
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
