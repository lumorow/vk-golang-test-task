package service

import "filmlib/server/internal/repository"

type Authorization interface {
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
