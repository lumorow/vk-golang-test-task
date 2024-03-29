package service

import (
	"crypto/sha1"
	"errors"
	"filmlib/server/internal/entity"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	UserRole string `json:"user_role"`
}

const (
	salt       = "sdf123123sadfsdf125346"
	signingKey = "eyJhbG#ciOiJIUz#I1NiJ9"
	tokenTTL   = 12 * time.Hour
)

func (s *AuthorizationService) CreateUser(user entity.User) (int, error) {
	if _, ok := s.roles[user.Role]; !ok {
		return 0, errors.New("unknown role")
	}

	user.Password = generatePasswordHash(user.Password)

	return s.Authorization.CreateUser(user)
}

func (s *AuthorizationService) GenerateToken(username, password string) (string, error) {
	user, err := s.Authorization.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Role,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthorizationService) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.UserRole, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
