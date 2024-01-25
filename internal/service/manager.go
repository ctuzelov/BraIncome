package service

import (
	"braincome/internal/models"
	"braincome/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserByEmail(email string, password string) (models.User, error)
}

type Authentication interface {
	CheckAuthority(user_type, role string) bool
	GetUserInfo(userId string) (models.User, error)
}

type Service struct {
	Authorization
	Authentication
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repo.Authorization),
		Authentication: NewUserInfo(repo.Authentication),
	}
}
