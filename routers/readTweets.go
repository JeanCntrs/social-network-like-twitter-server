package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
)

// ReadTweets : Lee los tweets
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
	}

	pag := int64(page)
	response, correct := db.ReadTweets(ID, pag)
	if correct == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
