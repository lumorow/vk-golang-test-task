package repository

import "fmt"

func (ap *ActorPostgres) DeleteActorById(actorId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", actorsTable)
	_, err := ap.db.Exec(query, actorId)
	if err != nil {
		return err
	}
	return nil
}
