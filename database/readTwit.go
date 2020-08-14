package database

import (
	"context"
	"log"
	"time"

	"github.com/ryukazari/twinstgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadTwit read twits from database
func ReadTwit(ID string, pagina int64) ([]*models.ResponseTwit, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("bd_twinstgo")
	col := db.Collection("twit")

	var results []*models.ResponseTwit

	condicion := bson.M{"userid": ID}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "date", Value: -1}}) // Sorted documents desc
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.ResponseTwit
		err := cursor.Decode(&registro)
		if err != nil {
			log.Fatal(err.Error())
			return results, false
		}
		results = append(results, &registro)
	}
	return results, true
}
