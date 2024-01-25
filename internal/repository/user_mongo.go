package repository

import (
	"braincome/internal/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserMongo struct {
	db *mongo.Client
}

func NewUserMongo(db *mongo.Client) *UserMongo {
	return &UserMongo{db: db}
}

func (r *UserMongo) GetUser(userID string) (models.User, error) {
	collection := r.db.Database("cluster0").Collection("user")

	// Your query conditions based on userID
	filter := bson.M{"user_id": userID}

	// Execute the query
	var user models.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	return user, nil
}

func (r *UserMongo) GetUserType(userId string) (string, error) {
	collection := r.db.Database("cluster0").Collection("user")

	// Your query conditions based on userId
	filter := bson.M{"user_id": userId}

	// Projection to get only the user_type field
	projection := bson.M{"user_type": 1}

	// Execute the query
	var result bson.M
	err := collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.New("user not found")
		}
		return "", err
	}

	userType, ok := result["user_type"].(string)
	if !ok {
		return "", errors.New("user_type not found or not a string")
	}

	return userType, nil
}
