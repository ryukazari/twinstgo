package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ryukazari/twinstgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ModifyRegister modified user profile
func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("bd_twinstgo")
	col := db.Collection("users")

	registro := make(map[string]interface{})
	if len(u.Name) > 0 {
		registro["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		registro["lastName"] = u.LastName
	}
	registro["birthday"] = u.Birthday
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		registro["biography"] = u.Biography
	}
	if len(u.Ubication) > 0 {
		registro["ubication"] = u.Ubication
	}
	if len(u.WebSite) > 0 {
		registro["webSite"] = u.WebSite
	}

	updateString := bson.M{
		"$set": registro,
	}

	objectID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objectID}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
