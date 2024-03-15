package service

import (
	"errors"
	"filmlib/server/internal/entity"
	"fmt"
	"time"
)

func validationCreateActor(actor entity.Actor) error {
	if err := checkActorSexCreate(actor.Sex); err != nil {
		return err
	}
	if err := checkActorBirthdayCreate(actor.Birthday); err != nil {
		return err
	}
	return nil
}

func validationCreateFilm(film entity.Film) error {
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

func checkActorSexCreate(sex string) error {
	availableSex := map[string]struct{}{"male": {}, "female": {}}
	if _, ok := availableSex[sex]; !ok {
		return fmt.Errorf("non-existent gender: %s: ", sex)
	}
	return nil
}

func checkActorBirthdayCreate(birthday string) error {
	formatBirthday, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return err
	}
	today := time.Now().Truncate(24 * time.Hour)
	ok := formatBirthday.Before(today)
	if !ok {
		return fmt.Errorf("incorrect birthday")
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
