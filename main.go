package main

import (
	"log"
	"github.com/JeanCntrs/social-network-like-twitter-server/handlers"
	"github.com/JeanCntrs/social-network-like-twitter-server/bd"
) 

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión")
		return
	}
	handlers.RouteHandlers()
}
