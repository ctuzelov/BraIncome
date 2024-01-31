package repository

import (
	"braincome/internal/models"
	"context"
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

	// Insert the user document into the MongoDB collection
	_, err := collection.InsertOne(ctx, m)
	if err != nil {
		return err
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

// func (r *AuthMongo) UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

// 	UserCollection := r.db.Database("cluster0").Collection("user")

// 	var updateObj primitive.D

// 	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
// 	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

// 	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: Updated_at})

// 	upsert := true
// 	filter := bson.M{"user_id": userId}
// 	opt := options.UpdateOptions{
// 		Upsert: &upsert,
// 	}

// 	_, err := UserCollection.UpdateOne(
// 		ctx,
// 		filter,
// 		bson.D{
// 			{Key: "$set", Value: updateObj},
// 		},
// 		&opt,
// 	)

// 	defer cancel()

// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// }

// func (r *AuthMongo) SetUserId(foundUser *models.User) {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	UserCollection := r.db.Database("cluster0").Collection("user")

// 	UserCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)

// 	defer cancel()
// }

// func (r *AuthMongo) FindEmail(email string) (models.User, error) {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	UserCollection := r.db.Database("cluster0").Collection("user")

// 	var foundUser models.User

// 	err := UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&foundUser)
// 	defer cancel()
// 	if err != nil {
// 		return models.User{}, err
// 	}

// 	if foundUser.Email == "" {
// 		return models.User{}, fmt.Errorf(validator.MsgUserNotFound)
// 	}

// 	return foundUser, nil
// }
