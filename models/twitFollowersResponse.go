package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//TwitFollowersResponse follower twit model response
type TwitFollowersResponse struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userid,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userrelationid,omitempty"`
	Twit           struct {
		Message string `bson:"message" json:"message,omitempty"`
		Date    string `bson:"date" json:"date,omitempty"`
		ID      string `bson:"_id" json:"_id,omitempty"`
	}
}
