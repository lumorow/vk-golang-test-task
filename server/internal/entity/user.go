package entity

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
