package service

import (
	"braincome/internal/models"
	"braincome/internal/repository"
)

type UserInfo struct {
	repo repository.Authentication
}

func NewUserInfo(repo repository.Authentication) *UserInfo {
	return &UserInfo{
		repo: repo,
	}
}

func (u *UserInfo) CheckAuthority(user_type, role string) bool {

	userType, err := u.repo.GetUserType(user_type)

	if userType != role || err != nil {
		return false
	}

	return true
}

func (u *UserInfo) GetUserInfo(userId string) (models.User, error) {
	return u.repo.GetUser(userId)
}
