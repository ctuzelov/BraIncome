package service

import (
	"braincome/internal/helper"
	"braincome/internal/models"
	"braincome/internal/repository"
	"fmt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserByEmail(email string, password string) (models.User, error) {
	var foundUser models.User
	foundUser, err := s.repo.FindEmail(email)
	if err != nil {
		return models.User{}, err
	}

	passwordIsValid, msg := helper.VerifyPassword(password, foundUser.Password)
	if !passwordIsValid {
		return models.User{}, fmt.Errorf(msg)
	}

	token, refreshToken, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.First_name, foundUser.Last_name, foundUser.User_type, foundUser.User_id)
	s.repo.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	s.repo.SetUserId(&foundUser)

	return foundUser, nil

}
