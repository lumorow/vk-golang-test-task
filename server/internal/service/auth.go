package service

import (
	"filmlib/server/internal/repository"
)

type AuthService struct {
	repoAuth repository.Authorization
}

func NewAuthService(repoAuth repository.Authorization) *AuthService {
	return &AuthService{repoAuth: repoAuth}
}
