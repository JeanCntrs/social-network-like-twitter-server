package db

import (
	"github.com/JeanCntrs/social-network-like-twitter-server/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin : Comprueba si la contraseña enviada coincide con la del usuario correspondiente
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserExist(email)
	if found == false {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
