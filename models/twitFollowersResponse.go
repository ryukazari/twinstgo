package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TwitFollowersResponse follower twit model response
type TwitFollowersResponse struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userId,omitempty"`
	UserFollowedID string             `bson:"userfollowedid" json:"userFollowedId,omitempty"`
	Twit           struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
