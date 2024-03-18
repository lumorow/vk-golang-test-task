package service

import (
	"errors"
	"filmlib/server/internal/entity"
)

func (s *FilmService) GetFilmsWithSort(sortType string, filmsId []int) ([]entity.Film, error) {
	if sortType == "" {
		sortType = "rating"
	}

	sortsExists := map[string]struct{}{"name": {}, "rating": {}, "release": {}}

	if _, ok := sortsExists[sortType]; !ok {
		return nil, errors.New("unknown sort type")
	}

	res, err := s.Film.GetFilmsWithSort(sortType, filmsId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
