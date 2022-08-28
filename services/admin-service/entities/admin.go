package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id" `
	IdString string             `json:"id" bson:"-"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
}
