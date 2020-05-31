package middleware

import (
	"net/http"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
)

// CheckDB : Comprueba el estado de la base de datos y de estar todo bien sigue el proceso, caso contrario, detiene la ejecución
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
