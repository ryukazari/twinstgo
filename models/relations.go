package models

//Relation model to relation between user and user followed
type Relation struct {
	UserID         string `bson:"userid" json:"userId"`
	UserFollowedID string `bson:"userfollowedid" json:"userFollowedId"`
}
