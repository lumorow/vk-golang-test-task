package repository

import (
	"errors"
	"filmlib/server/internal/entity"
	"fmt"
	"strings"
)

func (r *FilmRepository) GetFilmsWithSort(sortType string, filmsId []int) ([]entity.Film, error) {
	var films []entity.Film

	var trimFilmsId string
	if len(filmsId) > 0 {
		trimFilmsId = fmt.Sprintf("WHERE id IN (%s)", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(filmsId)), ","), "[]"))
	}
	var query string

	switch sortType {
	case "rating":
		query = fmt.Sprintf("SELECT name, description, release, rating FROM %s %s ORDER BY rating DESC", filmsTable, trimFilmsId)
	case "release":
		query = fmt.Sprintf("SELECT name, description, release, rating FROM %s %s ORDER BY release DESC", filmsTable, trimFilmsId)
	case "name":
		query = fmt.Sprintf("SELECT name, description, release, rating FROM %s %s ORDER BY name", filmsTable, trimFilmsId)
	default:
		return nil, errors.New("unknown sort type")
	}

	err := r.db.Select(&films, query)
	if err != nil {
		return nil, err
	}

	return films, nil
}
