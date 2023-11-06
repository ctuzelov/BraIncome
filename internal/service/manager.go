package service

import (
	"braincome/internal/models"
	"braincome/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserByEmail(email string, password string) (models.User, error)
}

type User interface {
}

type Service struct {
	Authorization
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
