package service

import "filmlib/server/internal/entity"

func (fs *FilmService) CreateFilm(film entity.Film) (int, error) {
	if err := film.Validate(); err != nil {
		return 0, err
	}

	id, err := fs.filmRepo.CreateFilm(film)
	if err != nil {
		return 0, err
	}

	return id, nil
}
