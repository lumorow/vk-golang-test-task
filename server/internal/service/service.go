package service

import (
	"filmlib/server/internal/entity"
	"filmlib/server/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Actor interface {
}

type Film interface {
}

type Service struct {
	Authorization
	Actor
	Film
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewAuthService(repo.Authorization),
		NewActorService(repo.Actor, repo.Film),
		NewFilmService(repo.Film),
	}
}
