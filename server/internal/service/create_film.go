package service

import "filmlib/server/internal/entity"

func (fs *FilmService) CreateFilm(film entity.Film) (int, error) {
	if err := validationCreateFilm(film); err != nil {
		return 0, err
	}
	return 0, nil
}
