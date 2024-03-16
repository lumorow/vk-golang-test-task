package service

func (fs *FilmService) DeleteFilmById(id int) error {
	err := fs.filmRepo.DeleteFilmById(id)
	if err != nil {
		return err
	}

	return nil
}
