package models

//Twit body to twit created
type Twit struct {
	Message string `bson:"message" json:"message"`
}
