package routers

import (
	"encoding/json"
	"net/http"

	"github.com/frangar97/goapitwitter/bd"
	"github.com/frangar97/goapitwitter/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Error en los datos"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(usuario.Email) == 0 {
		http.Error(w, "El email es necesario", http.StatusBadRequest)
		return
	}

	if len(usuario.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(usuario.Email)

	if encontrado {
		http.Error(w, "Usuario ya ha sido registrado", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertoRegistro(usuario)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
