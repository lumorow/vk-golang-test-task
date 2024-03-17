package repository

import "fmt"

func (r *Repository) DeleteFilmById(filmId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", filmsTable)
	_, err := r.db.Exec(query, filmId)
	if err != nil {
		return err
	}
	return nil
}
