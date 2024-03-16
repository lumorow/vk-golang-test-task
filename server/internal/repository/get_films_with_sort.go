package repository

import "filmlib/server/internal/entity"

func (fp *FilmPostgres) GetFilmsWithSort(sortMode string, filmsId []int) ([]entity.Film, error) {
	return nil, nil
}
