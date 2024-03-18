package service

func (s *FilmService) DeleteFilmById(id int) error {
	err := s.Film.DeleteFilmById(id)
	if err != nil {
		return err
	}

	return nil
}
