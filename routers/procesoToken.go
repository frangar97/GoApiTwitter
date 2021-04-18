package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/frangar97/goapitwitter/bd"
	"github.com/frangar97/goapitwitter/models"
)

var (
	Email     string
	IDUsuario string
)

//ProcesoToken extrae los valores del token
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Probando")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido")
	}

	return claims, false, "", err
}
