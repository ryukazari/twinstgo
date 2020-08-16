package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ryukazari/twinstgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadAllUsers read all users form database
func ReadAllUsers(ID string, pagina int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnect.Database("bd_twinstgo")
	col := db.Collection("users")

	var resultado []*models.User
	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSkip((pagina - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return resultado, false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return resultado, false
		}
		var r models.Relation
		r.UserID = ID
		r.UserFollowedID = s.ID.Hex()

		incluir = false
		encontrado, err = ExistRelation(r)
		// new: listar todos los usuarios que no sigo
		if tipo == "new" && !encontrado {
			incluir = true
		}

		// follow: listar todos los usuarios que sigo
		if tipo == "follow" && encontrado {
			incluir = true
		}

		// Si me estoy siguiendo a mi mismo?¿ set false
		if r.UserFollowedID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Ubication = ""
			s.Banner = ""
			s.Email = ""
			resultado = append(resultado, &s)
		}

	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}
	cursor.Close(ctx)
	return resultado, true
}
