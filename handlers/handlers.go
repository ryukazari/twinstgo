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
	router.HandleFunc("/twinst-go/api/auth/login", middleware.CheckDB(routers.Login)).Methods("POST")

	router.HandleFunc("/twinst-go/api/user/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/twinst-go/api/user/verperfil", middleware.CheckDB(middleware.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/twinst-go/api/user/modificarPerfil", middleware.CheckDB(middleware.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/twinst-go/api/user/subirAvatar", middleware.CheckDB(middleware.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/twinst-go/api/user/obtenerAvatar", middleware.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/twinst-go/api/user/subirBanner", middleware.CheckDB(middleware.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/twinst-go/api/user/obtenerBanner", middleware.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/twinst-go/api/twit/crear", middleware.CheckDB(middleware.ValidateJWT(routers.CreateTwit))).Methods("POST")
	router.HandleFunc("/twinst-go/api/twit/leerTwits", middleware.CheckDB(middleware.ValidateJWT(routers.ReadTwits))).Methods("GET")
	router.HandleFunc("/twinst-go/api/twit/eliminarTwit", middleware.CheckDB(middleware.ValidateJWT(routers.DeleteTwit))).Methods("DELETE")

	router.HandleFunc("/twinst-go/api/relation/consulta", middleware.CheckDB(middleware.ValidateJWT(routers.ExistRelation))).Methods("GET")
	router.HandleFunc("/twinst-go/api/relation/crear", middleware.CheckDB(middleware.ValidateJWT(routers.CreateRelation))).Methods("POST")
	router.HandleFunc("/twinst-go/api/relation/eliminar", middleware.CheckDB(middleware.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Println("Server is running in port ", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
