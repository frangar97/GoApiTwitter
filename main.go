package main

import (
	"log"

	"github.com/frangar97/goapitwitter/bd"
	"github.com/frangar97/goapitwitter/handlers"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
