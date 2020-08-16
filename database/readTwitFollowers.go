package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ryukazari/twinstgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ReadTwitFollowers read the twits from the followers from database
func ReadTwitFollowers(ID string, pagina int) ([]models.TwitFollowersResponse, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("bd_twinstgo")
	col := db.Collection("relations")

	skip := (pagina - 1) * 20
	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"userid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "twit", // con que tabla se unirá la tabla relation
			"localField":   "userfollowedid",
			"foreignField": "userid",
			"as":           "twit",
		},
	})
	condiciones = append(condiciones, bson.M{"$unwind": "$twit"})          //
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"date": -1}}) //ordenamiento | -1: DESC | 1: ASC |
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.TwitFollowersResponse
	err = cursor.All(ctx, &result)
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}
	return result, true
}
