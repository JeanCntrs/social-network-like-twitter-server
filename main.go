package main

import (
	"log"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
	"github.com/JeanCntrs/social-network-like-twitter-server/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexión")
		return
	}
	handlers.RouteHandlers()
}
