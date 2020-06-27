package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
	"github.com/JeanCntrs/social-network-like-twitter-server/models"
)

// VerifyRelation : Chequea si hay relación entre 2 usuarios
func VerifyRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro id es obligatorio", http.StatusBadRequest)
		return
	}

	var rel models.Relation
	rel.UserID = UserID
	rel.UserRelationshipID = ID

	var resp models.ResponseConsultRelation

	status, err := db.ConsultRelation(rel)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
