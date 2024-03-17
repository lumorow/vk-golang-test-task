package entity

// User represents a user entity.
type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
	Role     string `json:"role" example:"admin"`
}
