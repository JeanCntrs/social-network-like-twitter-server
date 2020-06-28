package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
)

// ReadTweetsRelation : Lee los tweets de todos los seguidores
func ReadTweetsRelation(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	resp, correct := db.ReadTweetsFollowers(UserID, page)
	if correct == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
