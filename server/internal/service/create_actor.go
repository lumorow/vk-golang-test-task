package service

import "filmlib/server/internal/entity"

func (s *Service) CreateActor(actor entity.Actor) (int, error) {
	if err := actor.Validate(); err != nil {
		return 0, err
	}

	id, err := s.Repository.CreateActor(actor)
	if err != nil {
		return 0, err
	}

	return id, nil
}
