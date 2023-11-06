package repository

import (
	"braincome/internal/helper"
	"braincome/internal/models"
	"braincome/internal/validator"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthMongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user models.User) (int, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	UserCollection := r.db.Database("cluster0").Collection("user")

	ExistedEmails, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if ExistedEmails > 0 {
		return http.StatusInternalServerError, fmt.Errorf(validator.MsgEmailExists)
	}
	password := helper.HashPassword(user.Password)
	user.Password = password

	ExistedNumbers, err := UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	defer cancel()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if ExistedNumbers > 0 {
		return http.StatusInternalServerError, fmt.Errorf(validator.MsgNumberExists)
	}

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := helper.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.User_type, user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	resultInsertionNumber, insertErr := UserCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return http.StatusInternalServerError, insertErr
	}
	defer cancel()

	return resultInsertionNumber.InsertedID.(int), nil
}

func (r *AuthMongo) UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	UserCollection := r.db.Database("cluster0").Collection("user")

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: Updated_at})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := UserCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)

	defer cancel()

	if err != nil {
		log.Fatal(err)
		return
	}
}

func (r *AuthMongo) SetUserId(foundUser *models.User) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	UserCollection := r.db.Database("cluster0").Collection("user")

	UserCollection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)

	defer cancel()
}

func (r *AuthMongo) FindEmail(email string) (models.User, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	UserCollection := r.db.Database("cluster0").Collection("user")

	var foundUser models.User

	err := UserCollection.FindOne(ctx, bson.M{"email": email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return models.User{}, err
	}

	if foundUser.Email == "" {
		return models.User{}, fmt.Errorf(validator.MsgUserNotFound)
	}

	return foundUser, nil
}
