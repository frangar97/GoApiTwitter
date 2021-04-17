package models

//RespuestaLogin respuesta al hacer login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
