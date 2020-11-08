package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	todo "github.com/lvl0nax/tutorial_rest_todo"
	"github.com/lvl0nax/tutorial_rest_todo/pkg/repository"
)

const (
	salt       = "saf4a34f3feg56gf"
	signingKey = "RandSigningTokenBytes"
	tokenTTL   = 12 * time.Hour
)

type AuthSerivce struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthSerivce {
	return &AuthSerivce{repo}
}

func (s *AuthSerivce) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthSerivce) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	tokenStr, err := token.SignedString([]byte(signingKey))

	if nil != err {
		logrus.Errorln("Error while signing the token")
		logrus.Errorf("Error signing token: %v\n", err)
		return "", err
	}

	return tokenStr, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
