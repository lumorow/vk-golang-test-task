package entity

import (
	"fmt"
	"time"
)

type Actor struct {
	ID       int64  `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	Sex      string `json:"sex" db:"sex"`
	Birthday string `json:"birthday" db:"birthday"`
}

func (actor Actor) Validate() error {
	if err := checkActorSexCreate(actor.Sex); err != nil {
		return err
	}

	if err := checkActorBirthdayCreate(actor.Birthday); err != nil {
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
