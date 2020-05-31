package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
	"github.com/JeanCntrs/social-network-like-twitter-server/models"
)

// Register : Es la función para crear el registro de usuario en la base de datos
func Register(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nill {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
	}

	if len(u.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(u.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, found, _ := db.CheckUserExist(u.Email)
	if found == true {
		http.Error(w, "Ya existe un usuario registrado con este email", 400)
		return
	}

	_, status, err := db.InsertRegister(u)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
