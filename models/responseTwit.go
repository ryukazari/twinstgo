package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ResponseTwit struct for every twit to response to frontend
type ResponseTwit struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Message string             `bson:"message" json:"message"`
	Date    time.Time          `bson:"date" json:"date"`
}
