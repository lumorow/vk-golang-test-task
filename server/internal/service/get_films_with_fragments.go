package service

import (
	"filmlib/server/internal/entity"
)

func (s *Service) GetFilmWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error) {
	res, err := s.Repository.GetFilmsWithFragment(actorNameFrag, filmNameFrag)
	if err != nil {
		return nil, err
	}

	return res, nil
}
