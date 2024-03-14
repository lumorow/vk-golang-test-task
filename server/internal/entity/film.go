package entity

type Film struct {
	ID          int64  `json:"-" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	ReleaseDay  string `json:"releaseDay" db:"releaseDay"`
	Rating      int8   `json:"rating" db:"rating"`
}
