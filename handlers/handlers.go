package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JeanCntrs/social-network-like-twitter-server/middleware"
	"github.com/JeanCntrs/social-network-like-twitter-server/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// RouteHandlers : Levanta el servidor en un puerto especifico, dandole permisos a todos para consultar la API
func RouteHandlers() {
	// Mux : Captura el http y le da un manejo al request y response que viene en el llamado a la API
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/ver-perfil", middleware.CheckDB(middleware.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modificar-perfil", middleware.CheckDB(middleware.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/leer-tweets", middleware.CheckDB(middleware.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/eliminar-tweet", middleware.CheckDB(middleware.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/subir-avatar", middleware.CheckDB(middleware.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/subir-banner", middleware.CheckDB(middleware.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/obtener-avatar", middleware.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/obtener-banner", middleware.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/alta-relacion", middleware.CheckDB(middleware.ValidateJWT(routers.HighRelation))).Methods("POST")
	router.HandleFunc("/baja-relacion", middleware.CheckDB(middleware.ValidateJWT(routers.LowRelation))).Methods("DELETE")
	router.HandleFunc("/consulta-relacion", middleware.CheckDB(middleware.ValidateJWT(routers.VerifyRelation))).Methods("GET")

	router.HandleFunc("/lista-usuarios", middleware.CheckDB(middleware.ValidateJWT(routers.ListUsers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
