package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Jonagarme/twiiterBack/middlew"
	"github.com/Jonagarme/twiiterBack/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Funcion para llamar a la api*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods(("POST"))
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods(("POST"))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
