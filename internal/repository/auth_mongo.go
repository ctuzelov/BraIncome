package repository

import (
	"braincome/internal/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) InsertUser(m models.User) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")

	defer cancel()

	m.ID = primitive.NewObjectID()

	// Insert the user document into the MongoDB collection
	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthMongo) SetToken(id primitive.ObjectID, token string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")
	defer cancel()

	// Define the filter to update the user by ID
	filter := bson.M{"_id": id}

	// Define the update to set the token and expiration time
	update := bson.M{
		"$set": bson.M{
			"token":   token,
			"expires": time.Now().Add(8 * time.Hour),
		},
	}

	// Update the user in the MongoDB collection
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *AuthMongo) RemoveToken(token string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")
	defer cancel()

	// Define the filter to find the user by token
	filter := bson.M{"token": token}

	// Define the update to remove the token
	update := bson.M{"$set": bson.M{"token": nil}}

	// Update the user in the MongoDB collection
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *AuthMongo) SwitchRole(email string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")
	defer cancel()

	filter := bson.M{"email": email}

	update := bson.M{"$set": bson.M{"user_type": "ADMIN"}}

	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return fmt.Errorf("error in updating the role")
	}

	return nil
}

func (r *AuthMongo) UserByEmail(email string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")

	defer cancel()

	// Define the filter to find the user by email
	filter := bson.M{"email": email}

	// Find the user in the MongoDB collection
	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, models.ErrNoRecord
		}
		return user, err
	}

	return user, nil
}

func (r *AuthMongo) UserByName(name string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")
	defer cancel()

	// Define the filter to find the user by name
	filter := bson.M{"name": name}

	// Find the user in the MongoDB collection
	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, models.ErrNoRecord
		}
		return user, err
	}

	return user, nil
}

func (r *AuthMongo) UserByToken(token string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")
	defer cancel()

	// Define the filter to find the user by token and not expired
	filter := bson.M{"token": token, "expires": bson.M{"$gt": time.Now()}}

	// Find the user in the MongoDB collection
	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, models.ErrNoRecord
		}
		return user, err
	}

	return user, nil
}

func (r *AuthMongo) UserById(id primitive.ObjectID) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")
	defer cancel()

	// Define the filter to find the user by ID
	filter := bson.M{"_id": id}

	// Find the user in the MongoDB collection
	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, models.ErrNoRecord
		}
		return user, err
	}

	return user, nil
}

func (r *AuthMongo) UpdateTokens(signedToken string, signedRefreshToken string, user_type string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")
	defer cancel()

	// Define the filter to find the user by token
	filter := bson.M{"token": signedToken}

	// Define the update to set the new tokens
	update := bson.M{"$set": bson.M{"token": signedToken, "refresh_token": signedRefreshToken, "user_type": user_type}}

	// Update the user in the MongoDB collection
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *AuthMongo) DeleteTokens(email string) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := r.db.Database("cluster0").Collection("user")
	defer cancel()

	// Define the filter to find the user by email
	filter := bson.M{"email": email}

	// Define the update to unset the token and refresh_token fields
	update := bson.M{"$unset": bson.M{"token": "", "refresh_token": ""}}

	// Update the user's document to remove the token and refresh_token fields
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
