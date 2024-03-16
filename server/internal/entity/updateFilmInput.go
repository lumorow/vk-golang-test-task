package entity

import "errors"

// UpdateFilmInput represents input data for updating a film entity.
type UpdateFilmInput struct {
	Name        *string `json:"name,omitempty" db:"name" example:"Inception"`
	Description *string `json:"description,omitempty" db:"description" example:"New description"`
	ReleaseDay  *string `json:"releaseDay,omitempty" db:"releaseDay" example:"2010-07-16"`
	Rating      *int8   `json:"rating,omitempty" db:"rating" minimum:"0" maximum:"10" example:"6"`
}

func (i UpdateFilmInput) Validate() error {
	if i.Name == nil && i.Description == nil && i.ReleaseDay == nil && i.Rating == nil {
		return errors.New("update film has no values")
	}
	if i.Name != nil {
		if err := checkFilmNameCreate(*i.Name); err != nil {
			return err
		}
	}
	if i.Description != nil {
		if err := checkFilmDescriptionCreate(*i.Description); err != nil {
			return err
		}
	}
	if i.Rating != nil {
		if err := checkFilmRatingCreate(*i.Rating); err != nil {
			return err
		}
	}
	if i.ReleaseDay != nil {
		if err := checkFilmReleaseDayCreate(*i.ReleaseDay); err != nil {
			return err
		}
	}
	return nil
}
