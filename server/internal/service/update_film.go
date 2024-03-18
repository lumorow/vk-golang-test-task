package service

import "filmlib/server/internal/entity"

func (s *FilmService) UpdateFilmById(filmId int, film entity.UpdateFilmInput) error {
	if err := film.Validate(); err != nil {
		return err
	}

	deleteIds := make([]int, 0)
	addIds := make([]int, 0)

	if film.ActorsId != nil {
		newActorsId := make(map[int]struct{}, len(*film.ActorsId))
		for _, newActorId := range *film.ActorsId {
			newActorsId[newActorId] = struct{}{}
		}

		actorsId, err := s.Actor.GetActorsIdByFilmId(filmId)
		if err != nil {
			return err
		}

		existingActorsId := make(map[int]struct{}, len(actorsId))
		for _, existingActorId := range actorsId {
			existingActorsId[existingActorId] = struct{}{}
		}

		deleteIds = make([]int, 0, len(actorsId))
		for _, actorId := range actorsId {
			if _, ok := newActorsId[actorId]; !ok {
				deleteIds = append(deleteIds, actorId)
			}
		}

		addIds = make([]int, 0, len(*film.ActorsId))
		for _, actorId := range *film.ActorsId {
			if _, ok := existingActorsId[actorId]; !ok {
				addIds = append(addIds, actorId)
			}
		}
	}

	err := s.Film.UpdateFilmById(filmId, deleteIds, addIds, film)
	if err != nil {
		return err
	}

	return nil
}
