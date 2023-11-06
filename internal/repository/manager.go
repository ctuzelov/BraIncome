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

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
	}
}
