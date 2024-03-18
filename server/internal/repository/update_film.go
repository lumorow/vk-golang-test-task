package repository

import (
	"filmlib/server/internal/entity"
	"fmt"
	"strings"
)

func (r *FilmRepository) UpdateFilmById(filmId int, deleteIds []int, addIds []int, film entity.UpdateFilmInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, deleteId := range deleteIds {
		queryDeleteIds := fmt.Sprintf("DELETE FROM %s WHERE actor_id = $1 AND film_id = $2", actorsFilmsTable)
		_, err = tx.Exec(queryDeleteIds, deleteId, filmId)
		if err != nil {
			return tx.Rollback()
		}
	}

	for _, addId := range addIds {
		createActorsFilmsIdsQuery := fmt.Sprintf("INSERT INTO %s (actor_id, film_id) values ($1, $2)", actorsFilmsTable)
		_, err = tx.Exec(createActorsFilmsIdsQuery, addId, filmId)
		if err != nil {
			return tx.Rollback()
		}
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if film.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *film.Name)
		argId++
	}

	if film.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *film.Description)
		argId++
	}

	if film.ReleaseDay != nil {
		setValues = append(setValues, fmt.Sprintf("release=$%d", argId))
		args = append(args, *film.ReleaseDay)
		argId++
	}

	if film.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *film.Rating)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", filmsTable, setQuery, argId)

	args = append(args, filmId)

	_, err = tx.Exec(query, args...)
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
