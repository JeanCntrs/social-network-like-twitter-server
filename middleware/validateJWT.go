package middleware

import (
	"net/http"

	"github.com/JeanCntrs/social-network-like-twitter-server/routers"
)

// ValidateJWT : Valida el JWT que viene en la petición
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error al procesar token"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
