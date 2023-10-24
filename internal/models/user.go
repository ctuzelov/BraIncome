package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type LoginType struct {
	ID   int    `bson:"id"`
	Name string `bson:"name"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id "`
	GUID        string             `bson:"guid" json:"guid"`
	Login       string             `bson:"login" json:"login"`
	LoginType   LoginType          `bson:"logintype" json:"logintype"`
	Name        string             `bson:"name" json:"name"`
	LastName    string             `bson:"lastname" json:"lastname"`
	LastLoginAt string             `bson:"lastloginat" json:"lastloginat"`
	CreatedAt   string             `bson:"createdat" json:"createdat"`
}
