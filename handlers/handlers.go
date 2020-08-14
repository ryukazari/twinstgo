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
	router.HandleFunc("/twinst-go/api/auth/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/twinst-go/api/user/verperfil", middleware.CheckDB(middleware.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/twinst-go/api/user/modificarPerfil", middleware.CheckDB(middleware.ValidateJWT(routers.ModifyProfile))).Methods("PUT")

	PORT := os.Getenv("TWINSTGO_PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Println("Server is running in port ", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
