package service

import (
	"filmlib/server/internal/entity"
)

//go:generate mockgen -destination=mocks/handler.go -package=mock -source=command_handler.go
//go:generate touch mocks/.coverignore

type Repository interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
	CreateActor(actor entity.Actor) (int, error)
	DeleteActorById(actorId int) error
	UpdateActorById(actorId int, actor entity.UpdateActorInput) error
	GetActor(actorId int) (entity.Actor, error)
	GetActorsIdByFilmId(filmId int) ([]int, error)
	CreateFilm(film entity.Film) (int, error)
	DeleteFilmById(filmId int) error
	GetFilmsByActorId(actorId int) ([]entity.Film, error)
	GetFilmsWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error)
	GetFilmsWithSort(sortType string, filmsId []int) ([]entity.Film, error)
	UpdateFilmById(filmId int, deleteIds []int, addIds []int, film entity.UpdateFilmInput) error
}

type Service struct {
	Repository
	roles map[string]struct{}
}

func NewService(repository Repository) *Service {
	roles := map[string]struct{}{"admin": {}, "user": {}}
	return &Service{
		repository,
		roles,
	}
}
