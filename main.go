package main

import (
	"log"

	"github.com/Jonagarme/twiiterBack/bd"
	"github.com/Jonagarme/twiiterBack/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
} // end main
