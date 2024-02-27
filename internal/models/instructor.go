package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Instructor struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	First_name   string             `bson:"first_name" json:"first_name"`
	Last_name    string             `bson:"last_name" json:"last_name"`
	Email        string             `bson:"email" json:"email"`
	Speciality   string             `bson:"speciality" json:"speciality"`
	Rating       float64            `bson:"rating" json:"rating"`
	Courses      int                `bson:"courses" json:"courses"`
	TotalReviews int                `bson:"total_reviews" json:"total_reviews"`
	About        string             `bson:"about" json:"about"`
	AvatarLink   string             `bson:"avatar_link" json:"avatar_link"`
}
