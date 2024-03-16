package entity

import (
	"errors"
)

// UpdateActorInput represents input data for updating an actor entity.
type UpdateActorInput struct {
	Name     *string `json:"name,omitempty" db:"name" example:"John Doe"`
	Sex      *string `json:"sex,omitempty" db:"sex" example:"female"`
	Birthday *string `json:"birthday,omitempty" db:"birthday" example:"1992-12-12"`
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
