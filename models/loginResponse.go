package models

// LoginResponse : Contiene el token devuelto al iniciar de sesión
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
