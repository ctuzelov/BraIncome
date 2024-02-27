package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title          string             `bson:"title" json:"title"`
	Rating         float64            `bson:"rating" json:"rating"`
	Price          int32              `bson:"price" json:"price"`
	Reviews        int                `bson:"reviews" json:"reviews"`
	Enrolled       int                `bson:"enrolled" json:"enrolled"`
	LastUpdated    time.Time          `bson:"last_updated" json:"last_updated"`
	Language       string             `bson:"language" json:"language"`
	Curriculum     []Module           `bson:"curriculum" json:"curriculum"` // References to Module collection
	Description    string             `bson:"description" json:"description"`
	Instructor     Instructor         `bson:"instructor" json:"instructor"`
	ReviewsList    []Review           `bson:"reviews_list" json:"reviews_list"`
	Categories     []string           `bson:"categories" json:"categories"`
	CoverPhotoLink string             `bson:"cover_photo_link" json:"cover_photo_link"`
}

type Lesson struct {
	Name string `form:"name"`
	Link string `form:"link"`
}

type Module struct {
	Name    string   `form:"name"`
	Lessons []Lesson `form:"lessons[]"`
}

type Review struct {
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	UserName string             `bson:"user_name" json:"user_name"`
	Rating   float64            `bson:"rating" json:"rating"`
	Comment  string             `bson:"comment" json:"comment"`
	Date     string             `bson:"date" json:"date"`
}
