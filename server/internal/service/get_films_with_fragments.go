package service

import (
	"filmlib/server/internal/entity"
)

func (fs *FilmService) GetFilmWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error) {
	res, err := fs.filmRepo.GetFilmsWithFragment(actorNameFrag, filmNameFrag)
	if err != nil {
		return nil, err
	}

	return res, nil
}
