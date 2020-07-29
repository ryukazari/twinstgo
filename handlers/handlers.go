package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/ryukazari/twinstgo/middleware"
	"github.com/ryukazari/twinstgo/routers"
)

// Handlers set port, handler and run server
func Handlers() {
	router := mux.NewRouter()

	//routes
	router.HandleFunc("/twinst-go/api/user/register", middleware.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("TWINSTGO_PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
