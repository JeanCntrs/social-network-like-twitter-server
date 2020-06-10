package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
	"github.com/JeanCntrs/social-network-like-twitter-server/models"
)

// ModifyProfile : Modifica el perfil de usuario
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var u models.User
	var status bool

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	status, err = db.ModifyRegister(u, UserID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro de usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
