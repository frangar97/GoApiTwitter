package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/frangar97/goapitwitter/bd"
	"github.com/frangar97/goapitwitter/jwt"
	"github.com/frangar97/goapitwitter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "Email de usuario es requerido", http.StatusBadRequest)
	}

	usuarioBD, existe := bd.IntentoLogin(usuario.Email, usuario.Password)

	if !existe {
		http.Error(w, "Usuario y/o Contraseña invalidos", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GeneroJWT(usuarioBD)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token", http.StatusBadRequest)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
