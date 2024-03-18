package service

import (
	"filmlib/server/internal/entity"
)

func (s *FilmService) GetFilmsWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error) {
	res, err := s.Film.GetFilmsWithFragment(actorNameFrag, filmNameFrag)
	if err != nil {
		return nil, err
	}

	return res, nil
}
