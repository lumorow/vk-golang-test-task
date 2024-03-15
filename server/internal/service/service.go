package service

import (
	"filmlib/server/internal/entity"
	"filmlib/server/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
	GetFilmWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error)
	GetFilmsWithSort(sortMode string) ([]entity.Film, error)
}

type Service struct {
	Authorization
	Actor
	Film
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewAuthService(repo.Authorization),
		NewActorService(repo.Actor, repo.Film),
		NewFilmService(repo.Actor, repo.Film),
	}
}
