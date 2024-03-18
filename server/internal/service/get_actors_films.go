package service

import (
	"filmlib/server/internal/entity"
)

func (s *ActorService) GetActorsWithFilms(actorsId []int) ([]entity.ActorFilms, error) {
	actorsFilms := make([]entity.ActorFilms, 0, len(actorsId))
	for i := 0; i < len(actorsId); i++ {
		var actorFilms entity.ActorFilms
		actorFilms.Films = make([]entity.Film, 0)

		actor, err := s.Actor.GetActor(actorsId[i])
		if err != nil {
			return nil, err
		}

		actorFilms.Name = actor.Name
		actorFilms.Sex = actor.Sex
		actorFilms.Birthday = actor.Birthday

		films, err := s.Film.GetFilmsByActorId(actorsId[i])
		if err != nil {
			return nil, err
		}

		actorFilms.Films = append(actorFilms.Films, films...)

		actorsFilms = append(actorsFilms, actorFilms)
	}

	return actorsFilms, nil
}
