package service

import (
	"braincome/internal/hasher"
	"braincome/internal/models"
	"braincome/internal/repository"
	"braincome/util"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repo repository.Authentication) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (u *AuthService) SignUp(m models.User) error {
	user, err := u.repo.UserByEmail(m.Email)

	if err != nil && !errors.Is(err, models.ErrNoRecord) {
		return fmt.Errorf("user_info - signup #1: %w", err)
	}
	if user.Email == m.Email {
		return models.ErrDuplicateEmail
	}
	user, err = u.repo.UserByName(m.First_name)
	if err != nil && !errors.Is(err, models.ErrNoRecord) {
		return err
	}
	// if user.First_name == m.First_name {
	// 	return models.ErrDuplicateName
	// }
	m.Password, err = hasher.Encrypt(m.Password)
	if err != nil {
		return err
	}

	m.Token, m.Refresh_Token, err = util.GenerateAllTokens(m.Email, m.First_name, m.User_type, (util.ObjectIDToString(user.ID)))

	if err != nil {
		return err
	}
	u.repo.InsertUser(m)
	return nil
}

func (u *AuthService) SignIn(login, password string) (models.User, error) {
	m, err := u.repo.UserByEmail(login)
	switch {
	case errors.Is(err, models.ErrNoRecord):
		m, err = u.repo.UserByName(login)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				return m, err
			}
			return m, err
		}
	case err != nil:
		return m, err
	}

	if !hasher.CorrectPassword(m.Password, password) {
		return m, models.ErrInvalidCredentials
	}

	m.Token, m.Refresh_Token, err = util.GenerateAllTokens(m.Email, m.First_name, m.User_type, util.ObjectIDToString(m.ID))
	if err != nil {
		return m, err
	}

	u.UpdateAllTokens(m.Token, m.Refresh_Token, m.User_type, m.ID)

	return m, nil
}

func (a *AuthService) UpdateAllTokens(signedToken string, signedRefreshToken string, user_type string, id primitive.ObjectID) (string, string, error) {
	const op = "util.UpdateAllTokens"

	_, err := util.ValidateToken(signedRefreshToken)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	err = a.repo.UpdateTokens(signedToken, signedRefreshToken, user_type)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return signedToken, signedRefreshToken, nil
}

func (u *AuthService) DeleteTokensByEmail(email string) error {
	const op = "service.user.DeleteTokensByEmail"

	err := u.repo.DeleteTokens(email)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (u *AuthService) GetByToken(token string) (models.User, error) {
	return u.repo.UserByToken(token)
}

func (u *AuthService) LogOut(token string) error {
	return u.repo.RemoveToken(token)
}

func (u *AuthService) MakeAdmin(email string) error {
	err := u.repo.SwitchRole(email)
	return err
}
