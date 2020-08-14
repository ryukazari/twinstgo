package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ryukazari/twinstgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ExistRelation resolve true if relation exists
func ExistRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("bd_twinstgo")
	col := db.Collection("relations")

	condicion := bson.M{
		"userid":         t.UserID,
		"userfollowedid": t.UserFollowedID,
	}
	var resultado models.Relation
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
