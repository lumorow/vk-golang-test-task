package service

func (s *ActorService) DeleteActorById(id int) error {
	err := s.Actor.DeleteActorById(id)
	if err != nil {
		return err
	}

	return nil
}
