package main

import (
	"log"

	"github.com/ryukazari/twinstgo/database"
	"github.com/ryukazari/twinstgo/handlers"
)

func main() {
	if !database.CheckConnection() {
		log.Fatal("Failed to connect to the database")
		return
	}
	handlers.Handlers()
}
