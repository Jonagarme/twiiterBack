package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jonagarme/twiiterBack/bd"
	"github.com/Jonagarme/twiiterBack/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}

	if len(t.Password) <= 6 {
		http.Error(w, "Debe especificar una contraseña min 6 digitos", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de user"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "NO se ha logrado insertar el registro de user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
