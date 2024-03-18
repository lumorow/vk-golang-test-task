package service

import (
	"filmlib/server/internal/entity"
)

//go:generate mockgen -destination=mocks/service.go -package=mock -source=service.go
//go:generate touch mocks/.coverignore

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

type AuthorizationService struct {
	Authorization
	roles map[string]struct{}
}

type ActorService struct {
	Actor
	Film
}

type FilmService struct {
	Actor
	Film
}

func NewAuthorizationService(authRepository Authorization) *AuthorizationService {
	roles := map[string]struct{}{"admin": {}, "user": {}}
	return &AuthorizationService{
		authRepository,
		roles,
	}
}

func NewActorService(actorRepository Actor, filmRepository Film) *ActorService {
	return &ActorService{
		actorRepository,
		filmRepository,
	}
}

func NewFilmService(actorRepository Actor, filmRepository Film) *FilmService {
	return &FilmService{
		actorRepository,
		filmRepository,
	}
}
