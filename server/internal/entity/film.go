package entity

import (
	"errors"
	"fmt"
	"time"
)

type Film struct {
	ID          int64  `json:"-" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	ReleaseDay  string `json:"releaseDay" db:"releaseDay"`
	Rating      int8   `json:"rating" db:"rating"`
}

func (film Film) Validate() error {
	if err := checkFilmNameCreate(film.Name); err != nil {
		return err
	}

	if err := checkFilmDescriptionCreate(film.Description); err != nil {
		return err
	}

	if err := checkFilmRatingCreate(film.Rating); err != nil {
		return err
	}

	if err := checkFilmReleaseDayCreate(film.ReleaseDay); err != nil {
		return err
	}

	return nil
}

func checkFilmNameCreate(name string) error {
	l := len(name)

	if l < 1 {
		return errors.New("minimum film title length = 1")
	}

	if l > 150 {
		return fmt.Errorf("maximum film title length = 150, your movie title length = %d", l)
	}

	return nil
}

func checkFilmDescriptionCreate(description string) error {
	l := len(description)

	if l > 1000 {
		return fmt.Errorf("maximum film description length = 1000, your movie title length = %d", l)
	}

	return nil
}

func checkFilmRatingCreate(rating int8) error {
	if rating < 0 || rating > 10 {
		return errors.New("film rating range from 0 to 10")
	}

	return nil
}

func checkFilmReleaseDayCreate(releaseDay string) error {
	_, err := time.Parse("2006-01-02", releaseDay)
	if err != nil {
		return err
	}

	return nil
}
