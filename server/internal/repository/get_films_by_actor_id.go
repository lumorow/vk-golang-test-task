package repository

import "filmlib/server/internal/entity"

func (fp *FilmPostgres) GetFilmsByActorId(actorId int) ([]entity.Film, error) {
	return nil, nil
}
