package service

import (
	"filmlib/server/internal/entity"
)

func (as *ActorService) GetActorsWithFilms(actorsId []int) ([]entity.ActorFilms, error) {
	res, err := as.actorRepo.GetActorsWithFilms(actorsId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
