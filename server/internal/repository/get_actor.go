package repository

import (
	"filmlib/server/internal/entity"
	"fmt"
)

func (r *Repository) GetActor(actorId int) (entity.Actor, error) {
	var actor entity.Actor

	query := fmt.Sprintf("SELECT name, sex, birthday FROM %s WHERE id = $1", actorsTable)

	if err := r.db.Get(&actor, query, actorId); err != nil {
		return actor, err
	}

	return actor, nil
}
