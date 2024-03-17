package repository

import (
	"filmlib/server/internal/entity"
	"fmt"
)

func (fp *FilmPostgres) GetFilmsByActorId(actorId int) ([]entity.Film, error) {
	var filmsByActorIf []entity.Film
	query := fmt.Sprintf(`SELECT f.name, f.description, f.release, f.release FROM films f
    											INNER JOIN actors_films af ON f.id = af.film_id WHERE af.actor_id = $1`)
	if err := fp.db.Select(&filmsByActorIf, query, actorId); err != nil {
		return nil, err
	}
	return filmsByActorIf, nil
}
