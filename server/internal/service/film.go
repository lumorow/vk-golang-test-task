package service

import "filmlib/server/internal/repository"

type FilmService struct {
	filmRepo repository.Film
}

func NewFilmService(filmRepo repository.Film) Film {
	return FilmService{
		filmRepo,
	}
}
