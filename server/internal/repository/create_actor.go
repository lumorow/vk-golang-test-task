package repository

import (
	"filmlib/server/internal/entity"
	"fmt"
)

func (r *ActorRepository) CreateActor(actor entity.Actor) (int, error) {
	var actorId int
	createActorQuery := fmt.Sprintf("INSERT INTO %s (name, sex, birthday) values ($1, $2, $3) RETURNING id", actorsTable)

	row := r.db.QueryRow(createActorQuery, actor.Name, actor.Sex, actor.Birthday)
	if err := row.Scan(&actorId); err != nil {
		return 0, err
	}

	return actorId, nil
}
