package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/frangar97/goapitwitter/bd"
	"github.com/frangar97/goapitwitter/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "Error en los datos", http.StatusBadRequest)
		return
	}

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el tweet.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
