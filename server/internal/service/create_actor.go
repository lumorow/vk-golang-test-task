package service

import "filmlib/server/internal/entity"

func (as *ActorService) CreateActor(actor entity.Actor) (int, error) {
	if err := validationCreateActor(actor); err != nil {
		return 0, err
	}
	return 0, nil
}
