package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title       string               `bson:"title" json:"title"`
	Rating      float64              `bson:"rating" json:"rating"`
	Reviews     int                  `bson:"reviews" json:"reviews"`
	Enrolled    int                  `bson:"enrolled" json:"enrolled"`
	CreatedBy   string               `bson:"created_by" json:"created_by"`
	LastUpdated time.Time            `bson:"last_updated" json:"last_updated"`
	Language    string               `bson:"language" json:"language"`
	Curriculum  []primitive.ObjectID `bson:"curriculum" json:"curriculum"` // References to Module collection
	Description string               `bson:"description" json:"description"`
	Instructor  Instructor           `bson:"instructor" json:"instructor"`
	ReviewsList []primitive.ObjectID `bson:"reviews" json:"reviews_list"`
	Categories  []string             `bson:"categories" json:"categories"`
}

type Module struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title   string             `bson:"title" json:"title"`
	Lessons []primitive.ObjectID           `bson:"lessons" json:"lessons"`
}

type Lesson struct {
	Title    string `bson:"title" json:"title"`
	Link     string `bson:"link" json:"link"`
	Duration string `bson:"duration" json:"duration"`
}

type Instructor struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Role         string             `bson:"role" json:"role"`
	Rating       float64            `bson:"rating" json:"rating"`
	Courses      int                `bson:"courses" json:"courses"`
	TotalReviews int                `bson:"total_reviews" json:"total_reviews"`
	About        string             `bson:"about" json:"about"`
}

type Review struct {
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	UserName string             `bson:"user_name" json:"user_name"`
	Rating   float64            `bson:"rating" json:"rating"`
	Comment  string             `bson:"comment" json:"comment"`
	Date     string             `bson:"date" json:"date"`
}
