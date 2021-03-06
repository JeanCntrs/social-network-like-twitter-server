package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/JeanCntrs/social-network-like-twitter-server/db"
	"github.com/JeanCntrs/social-network-like-twitter-server/models"
)

// UploadAvatar : Sube el avatar al servidor
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var newFileName string = "uploads/avatars/" + UserID + "." + extension

	f, err := os.OpenFile(newFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = UserID + "." + extension
	status, err = db.ModifyRegister(user, UserID)
	if err != nil || status == false {
		http.Error(w, "Error al grabar avatar en base de datos! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
