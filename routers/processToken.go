package routers

import (
	"errors"
	"strings"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email : Valor de email usado en todos los endpoints
var Email string

// UserID : Es el ID devuelto del modelo, se usará en todos los endpoints
var UserID string

// ProcessToken : Procesa el token y extrae sus valores
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myPassword := []byte("P4ssw0rd t0 s1gn t0k3n AxZ 938")
	claims := &models.Claim()

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token inválido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.token) (interface{}, error) {
		return myPassword, nil
	})
	if err == nil {
		_, found, _ := db.CheckUserExist(claims.Email)
		if found == true {
			Email = claims.Email
			UserId = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}

	return claims, false, string(""), err
}
