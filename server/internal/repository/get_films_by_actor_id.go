package repository

import (
	"filmlib/server/internal/entity"
)

func (r *FilmRepository) GetFilmsByActorId(actorId int) ([]entity.Film, error) {
	var filmsByActorIf []entity.Film
	query := `SELECT f.name, f.description, f.release, f.release FROM films f
    											INNER JOIN actors_films af ON f.id = af.film_id WHERE af.actor_id = $1`
	if err := r.db.Select(&filmsByActorIf, query, actorId); err != nil {
		return nil, err
	}
	return filmsByActorIf, nil
}
