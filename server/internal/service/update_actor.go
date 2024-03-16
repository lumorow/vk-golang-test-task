package service

import "filmlib/server/internal/entity"

func (as *ActorService) UpdateActorById(id int, actor entity.UpdateActorInput) error {
	if err := actor.Validate(); err != nil {
		return err
	}

	err := as.actorRepo.UpdateActorById(id, actor)
	if err != nil {
		return err
	}

	return nil
}
