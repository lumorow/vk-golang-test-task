package repository

import (
	"fmt"
)

func (r *ActorRepository) GetActorsIdByFilmId(filmId int) ([]int, error) {
	query := fmt.Sprintf("SELECT actor_id FROM %s WHERE film_id = $1", actorsFilmsTable)
	rows, err := r.db.Query(query, filmId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	actorsId := make([]int, 0)

	for rows.Next() {
		var actorID int
		err := rows.Scan(&actorID)
		if err != nil {
			return nil, err
		}
		actorsId = append(actorsId, actorID)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return actorsId, nil
}
