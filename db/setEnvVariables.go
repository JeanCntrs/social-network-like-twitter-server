package db

import (
	"os"
)

// setConnectionString : Setea la cadena de conexión
func setConnectionString() {
	os.Setenv("CONNECTION_STRING", "mongodb+srv://JeanCntrs:qCOTVU1QQpfmmcmY@cluster0-4ct1j.mongodb.net/SocialNetworkLikeTwitter")
}
