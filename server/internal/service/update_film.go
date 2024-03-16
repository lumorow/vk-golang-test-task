package service

import "filmlib/server/internal/entity"

func (fs *FilmService) UpdateFilmById(id int, film entity.UpdateFilmInput) error {
	if err := film.Validate(); err != nil {
		return err
	}

	err := fs.filmRepo.UpdateFilmById(id, film)
	if err != nil {
		return err
	}

	return nil
}
