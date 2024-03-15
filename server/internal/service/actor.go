package service

import (
	"filmlib/server/internal/repository"
)

type ActorService struct {
	actorRepo repository.Actor
	filmRepo  repository.Film
}

func NewActorService(actorRepo repository.Actor, filmRepo repository.Film) Actor {
	return &ActorService{
		actorRepo,
		filmRepo,
	}
}
