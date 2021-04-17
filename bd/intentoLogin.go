package bd

import (
	"github.com/frangar97/goapitwitter/models"
	"golang.org/x/crypto/bcrypt"
)

//IntentoLogin verifica que las credenciales del usuario coincidan con algun registro de la base de datos
func IntentoLogin(email, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)

	if !encontrado {
		return usuario, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usuario, false
	}

	return usuario, true
}
