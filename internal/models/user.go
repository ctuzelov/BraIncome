package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name string             `json:"first_name" validate:"required, min=2, max=100"`
	Last_name  string             `json:"last_name" validate:"required, min=2, max=100"`
	Password   string             `json:"password" validate:"required,min=6"`
	Email      string             `json:"email" validate:"required,email"`
	Token      *string            `json:"token"`
	User_type  string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Expires    *time.Time         `bson:"expires" json:"expires"`
	// AccessibleVideos map[string]bool           `json:"accessible_videos"`
}
