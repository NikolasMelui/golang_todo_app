package service

import (
	"crypto/sha256"
	"fmt"

	todo "github.com/nikolasmelui/golang_todo_app"
	"github.com/nikolasmelui/golang_todo_app/pkg/repository"
	"github.com/spf13/viper"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = HashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func HashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(viper.GetString("security.salt"))))
}
