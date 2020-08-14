package database

import (
	"context"
	"time"

	"github.com/ryukazari/twinstgo/models"
)

//DeleteRelation delete a relation from database
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("bd_twinstgo")
	col := db.Collection("relations")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
