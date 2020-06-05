package jwt

import (
	"time"

	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT : Genera el encriptado con JWT
func GenerateJWT(u models.User) (string, error) {
	myPassword := []byte("P4ssw0rd t0 s1gn t0k3n AxZ 938")
	payload := jwt.MapClaims{
		"email":     u.Email,
		"names":     u.Names,
		"surnames":  u.Surnames,
		"birthdate": u.Birthdate,
		"biography": u.Biography,
		"location":  u.Location,
		"website":   u.Website,
		"_id":       u.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPassword)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
