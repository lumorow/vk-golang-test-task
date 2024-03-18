package repository

import (
	"filmlib/server/internal/entity"
	"fmt"

	"github.com/sirupsen/logrus"
)

func (r *AuthorizationRepository) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, role) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Username, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorizationRepository) GetUser(username, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id, role FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	logrus.Print(user)

	return user, err
}
