package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
)

// ListUsers : Leo la lista de los usuarios
func ListUsers(w http.ResponseWriter, r *http.Request) {
	userType := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	tempPage, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(tempPage)

	result, status := db.ReadAllUsers(UserID, pag, search, userType)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
