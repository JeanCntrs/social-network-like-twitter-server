package routers

import (
	"net/http"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
)

// DeleteTweet : Recibe la petición que elimina el tweet
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, UserID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar eliminar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
