package service

import (
	"braincome/internal/models"
	"braincome/internal/repository"
)

type User interface {
	SignUp(models.User) error
	SignIn(login, password string) (models.User, error)
	GetByToken(token string) (models.User, error)
	LogOut(token string) error
}

type Service struct {
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserInfo(repo.Authentication),
	}
}
