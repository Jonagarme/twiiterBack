package middlew

import (
	"net/http"

	"github.com/Jonagarme/twiiterBack/bd"
)

/*ChequeoBD es el middlew que permite conocer el estado dela BD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
