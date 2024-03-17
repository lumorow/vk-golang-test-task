package service

import "filmlib/server/internal/entity"

func (s *Service) UpdateActorById(id int, actor entity.UpdateActorInput) error {
	if err := actor.Validate(); err != nil {
		return err
	}

	err := s.Repository.UpdateActorById(id, actor)
	if err != nil {
		return err
	}

	return nil
}
