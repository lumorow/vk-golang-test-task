package repository

import "fmt"

func (r *Repository) DeleteActorById(actorId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", actorsTable)
	_, err := r.db.Exec(query, actorId)
	if err != nil {
		return err
	}
	return nil
}
