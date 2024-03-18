package repository

import (
	"filmlib/server/internal/entity"
	"fmt"
)

func (r *FilmRepository) CreateFilm(film entity.Film) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var filmId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (name, description, release, rating) values ($1, $2, $3, $4) RETURNING id", filmsTable)

	row := tx.QueryRow(createItemQuery, film.Name, film.Description, film.ReleaseDay, film.Rating)
	err = row.Scan(&filmId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, actorId := range film.ActorsId {
		createActorsFilmsIdsQuery := fmt.Sprintf("INSERT INTO %s (actor_id, film_id) values ($1, $2)", actorsFilmsTable)
		_, err = tx.Exec(createActorsFilmsIdsQuery, actorId, filmId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	return filmId, tx.Commit()
}
