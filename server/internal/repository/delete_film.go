package repository

import "fmt"

func (fp *FilmPostgres) DeleteFilmById(filmId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", filmsTable)
	_, err := fp.db.Exec(query, filmId)
	if err != nil {
		return err
	}
	return nil
}
