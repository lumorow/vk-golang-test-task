package service

import "filmlib/server/internal/entity"

func (fs *FilmService) GetFilmsWithSort(sortMode string, filmsId []int) ([]entity.Film, error) {
	res, err := fs.filmRepo.GetFilmsWithSort(sortMode, filmsId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
