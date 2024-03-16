package service

func (as *ActorService) DeleteActorById(id int) error {
	err := as.actorRepo.DeleteActorById(id)
	if err != nil {
		return err
	}

	return nil
}
