package repository

import (
	"braincome/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user models.User) (int, error) {
	return 0, nil
}

func (r *AuthMongo) GetUser(username, password string) (models.User, error) {
	return models.User{}, nil
}
