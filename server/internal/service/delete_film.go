package service

func (s *Service) DeleteFilmById(id int) error {
	err := s.Repository.DeleteFilmById(id)
	if err != nil {
		return err
	}

	return nil
}
