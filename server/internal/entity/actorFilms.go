package entity

// ActorFilms represents an actor with their associated films.
type ActorFilms struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	Sex      string `json:"sex" db:"sex"`
	Birthday string `json:"birthday" db:"birthday"`
	Films    []Film
}
