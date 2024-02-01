package repository

import (
	"braincome/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authentication interface {
	InsertUser(models.User) error
	UserByEmail(email string) (models.User, error)
	UserByName(name string) (models.User, error)
	UserByToken(token string) (models.User, error)
	UserById(id primitive.ObjectID) (models.User, error)
	SetToken(id primitive.ObjectID, token string) error
	RemoveToken(token string) error
}

type Repository struct {
	Authentication
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authentication: NewAuthMongo(db),
	}
}
