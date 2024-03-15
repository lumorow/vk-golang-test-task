package service

import "filmlib/server/internal/repository"

type FilmService struct {
	actorRepo repository.Actor
	filmRepo  repository.Film
}

func NewFilmService(actorRepo repository.Actor, filmRepo repository.Film) Film {
	return &FilmService{
		actorRepo,
		filmRepo,
	}
}
