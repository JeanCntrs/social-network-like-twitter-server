package routers

import (
	"net/http"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
	"github.com/JeanCntrs/social-network-like-twitter-server/models"
)

// LowRelation : Elimina la relación entre usuarios
func LowRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro id es obligatorio", http.StatusBadRequest)
		return
	}

	var rel models.Relation
	rel.UserID = UserID
	rel.UserRelationshipID = ID

	status, err := db.DeleteRelation(rel)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar eliminar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado eliminar relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
