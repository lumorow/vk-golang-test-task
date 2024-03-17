package service

func (s *Service) DeleteActorById(id int) error {
	err := s.Repository.DeleteActorById(id)
	if err != nil {
		return err
	}

	return nil
}
