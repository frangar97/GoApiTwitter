package routers

import (
	"encoding/json"
	"net/http"

	"github.com/frangar97/goapitwitter/bd"
	"github.com/frangar97/goapitwitter/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
