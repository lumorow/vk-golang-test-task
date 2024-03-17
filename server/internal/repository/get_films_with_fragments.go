package repository

import (
	"filmlib/server/internal/entity"
	"fmt"
)

func (fp *FilmPostgres) GetFilmsWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error) {
	var films []entity.Film

	if actorNameFrag == "" {
		query := getFilmFragmentQuery()

		if err := fp.db.Select(&films, query, filmNameFrag); err != nil {
			return nil, err
		}
	} else if filmNameFrag == "" {
		query := getActorFragmentQuery()

		if err := fp.db.Select(&films, query, actorNameFrag); err != nil {
			return nil, err
		}
	} else {
		query := getActorFragmentFilmFragmentQuery()

		if err := fp.db.Select(&films, query, filmNameFrag, actorNameFrag); err != nil {
			return nil, err
		}
	}

	return films, nil
}

func getFilmFragmentQuery() string {
	return fmt.Sprintf(`SELECT DISTINCT f.name, f.description, f.release, f.rating FROM films f
              INNER JOIN %s af ON f.id = af.film_id
              INNER JOIN %s a ON af.actor_id = a.id
              WHERE f.name LIKE '%%' || $1 || '%%'`, actorsFilmsTable, actorsTable)
}

func getActorFragmentQuery() string {
	return fmt.Sprintf(`SELECT DISTINCT f.name, f.description, f.release, f.rating FROM films f
              INNER JOIN %s af ON f.id = af.film_id
              INNER JOIN %s a ON af.actor_id = a.id
              WHERE a.name LIKE '%%' || $1 || '%%'`, actorsFilmsTable, actorsTable)
}

func getActorFragmentFilmFragmentQuery() string {
	return fmt.Sprintf(`SELECT DISTINCT f.name, f.description, f.release, f.rating FROM films f
              INNER JOIN %s af ON f.id = af.film_id
              INNER JOIN %s a ON af.actor_id = a.id
              WHERE f.name LIKE '%%' || $1 || '%%' 
              OR a.name LIKE '%%' || $2 || '%%'`, actorsFilmsTable, actorsTable)
}
