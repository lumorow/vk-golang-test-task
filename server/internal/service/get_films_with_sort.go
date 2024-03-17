package service

import (
	"errors"
	"filmlib/server/internal/entity"
)

func (fs *FilmService) GetFilmsWithSort(sortType string, filmsId []int) ([]entity.Film, error) {
	if sortType == "" {
		sortType = "rating"
	}

	sortsEsxists := map[string]struct{}{"name": {}, "rating": {}, "release": {}}

	if _, ok := sortsEsxists[sortType]; !ok {
		return nil, errors.New("unknown sort type")
	}

	res, err := fs.filmRepo.GetFilmsWithSort(sortType, filmsId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
