package entity

import "errors"

type UpdateFilmInput struct {
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	ReleaseDay  *string `json:"releaseDay" db:"releaseDay"`
	Rating      *int8   `json:"rating" db:"rating"`
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
