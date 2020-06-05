package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
	"github.com/JeanCntrs/social-network-like-twitter-server/jwt"
	"github.com/JeanCntrs/social-network-like-twitter-server/models"
)

// Login : Comprueba si el usuario tiene acceso al sistema
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválida"+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	document, found := db.TryLogin(u.Email, u.Password)
	if found == false {
		http.Error(w, "Usuario y/o contraseña inválida", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el token"+err.Error(), 400)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
