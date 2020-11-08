package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/lvl0nax/tutorial_rest_todo"
	"github.com/lvl0nax/tutorial_rest_todo/pkg/repository"
)

const salt = "saf4a34f3feg56gf"

type AuthSerivce struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthSerivce {
	return &AuthSerivce{repo}
}

func (s *AuthSerivce) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
