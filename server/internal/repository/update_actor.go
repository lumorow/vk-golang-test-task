package repository

import (
	"filmlib/server/internal/entity"
	"fmt"
	"strings"
)

func (ap *ActorPostgres) UpdateActorById(actorId int, actor entity.UpdateActorInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if actor.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *actor.Name)
		argId++
	}

	if actor.Birthday != nil {
		setValues = append(setValues, fmt.Sprintf("birthday=$%d", argId))
		args = append(args, *actor.Birthday)
		argId++
	}

	if actor.Sex != nil {
		setValues = append(setValues, fmt.Sprintf("sex=$%d", argId))
		args = append(args, *actor.Sex)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id= $%d", actorsTable, setQuery, argId)

	args = append(args, actorId)

	_, err := ap.db.Exec(query, args)
	if err != nil {
		return err
	}
	return nil
}
