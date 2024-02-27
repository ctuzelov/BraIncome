package repository

import (
	"braincome/internal/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InstructorMongo struct {
	db *mongo.Client
}

func NewInstructorMongo(db *mongo.Client) *InstructorMongo {
	return &InstructorMongo{db: db}
}

func (i *InstructorMongo) Insert(instructor models.Instructor) error {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	collection := i.db.Database("cluster0").Collection("instructors")

	defer cancel()

	instructor.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, instructor)
	if err != nil {
		return err
	}

	return nil
}

func (i *InstructorMongo) Update(instructor models.Instructor) error {
	return nil
}

func (i *InstructorMongo) GetByName(name string) (models.Instructor, error) {
	collection := i.db.Database("cluster0").Collection("instructors")

	var instructor models.Instructor
	filter := bson.M{"name": name}
	err := collection.FindOne(context.Background(), filter).Decode(&instructor)
	if err != nil {
		return models.Instructor{}, models.ErrNoRecord
	}

	return instructor, nil
}

func (i *InstructorMongo) GetByEmail(email string) (models.Instructor, error) {
	collection := i.db.Database("cluster0").Collection("instructors")

	var instructor models.Instructor
	filter := bson.M{"email": email}
	err := collection.FindOne(context.Background(), filter).Decode(&instructor)
	if err != nil {
		return models.Instructor{}, models.ErrDuplicateEmail
	}

	return instructor, nil
}
