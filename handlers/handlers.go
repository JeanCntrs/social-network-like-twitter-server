package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// RouteHandlers : Levanta el servidor en un puerto especifico, dandole permisos a todos para consultar la API
func RouteHandlers()  {
	// Mux : Captura el http y le da un manejo al request y response que viene en el llamado a la API
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}