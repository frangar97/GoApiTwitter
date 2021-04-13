package middleware

import (
	"net/http"

	"github.com/frangar97/goapitwitter/bd"
)

//ChequeoBD es un middleware que verifica el estado de la base de datos
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConection() == 0 {
			http.Error(w, "No se pudo establecer conexión con base de datos", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	}
}
