package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/frangar97/goapitwitter/models"
)

func GeneroJWT(usuario models.Usuario) (string, error) {
	miClave := []byte("Probando")
	payload := jwt.MapClaims{
		"email":            usuario.Email,
		"nombre":           usuario.Nombre,
		"apellido":         usuario.Apellido,
		"fecha_nacimiento": usuario.FechaNacimiento,
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
