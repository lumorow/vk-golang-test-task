package entity

import (
	"errors"
)

type UpdateActorInput struct {
	Name     *string `json:"name" db:"name"`
	Sex      *string `json:"sex" db:"sex"`
	Birthday *string `json:"birthday" db:"birthday"`
}

func (i UpdateActorInput) Validate() error {
	if i.Name == nil && i.Sex == nil && i.Birthday == nil {
		return errors.New("update actor has no values")
	}
	if i.Sex != nil {
		if err := checkActorSexCreate(*i.Sex); err != nil {
			return err
		}
	}
	if i.Birthday != nil {
		if err := checkActorBirthdayCreate(*i.Birthday); err != nil {
			return err
		}
	}
	return nil
}
