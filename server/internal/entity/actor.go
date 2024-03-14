package entity

type Actor struct {
	ID       int64  `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	Sex      string `json:"sex" db:"sex"`
	Birthday string `json:"birthday" db:"birthday"`
}

type ActorFilms struct {
	ID       int64  `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	Sex      string `json:"sex" db:"sex"`
	Birthday string `json:"birthday" db:"birthday"`
	Films    []Film
}
