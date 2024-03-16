package entity

// User represents a user entity.
type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" example:"john_doe"`
	Password string `json:"password" example:"strongPassword"`
	Role     string `json:"role" example:"admin"`
}
