package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
	"github.com/JeanCntrs/social-network-like-twitter-server/models"
)

// SaveTweet : Permite grabar el tweet en la base de datos
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	registry := models.SaveTweet{
		UserID:    UserID,
		Message:   message.Message,
		CreatedAt: time.Now(),
	}

	_, status, err := db.InsertTweet(registry)
	if err != nil {
		http.Error(w, "Ocurrió un error al insertar el registro, reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se logró insertar el tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
