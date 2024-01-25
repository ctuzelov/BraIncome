package repository

import (
	"braincome/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	FindEmail(email string) (models.User, error)
	SetUserId(foundUser *models.User)
	UpdateAllTokens(signedToken string, signedRefreshToken string, userId string)
}

type Authentication interface {
	GetUser(userId string) (models.User, error)
	GetUserType(userId string) (string, error)
}

type Repository struct {
	Authorization
	Authentication
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization:  NewAuthMongo(db),
		Authentication: NewUserMongo(db),
	}
}
