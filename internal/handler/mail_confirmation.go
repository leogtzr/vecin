package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"vecin/internal/service"
)

func ConfirmAccountHandler(svc *service.Service, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]

	log.Printf("debug:x token: (%s)", token)
	err := svc.ConfirmAccount(token)
	if err != nil {
		log.Printf("error: %v", err)
		// TODO: redirect to error page
	}

	//var userID int
	//var expiry time.Time

	/*
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, "Error al iniciar la transacción", http.StatusInternalServerError)
			return
		}
		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				http.Error(w, "Error interno", http.StatusInternalServerError)
			}
		}()

		err = tx.QueryRow("SELECT usuario_id, fecha_expiracion FROM confirmacion_cuenta WHERE token = $1", token).Scan(&userID, &expiry)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Token inválido o expirado", http.StatusBadRequest)
			return
		}

		if time.Now().After(expiry) {
			tx.Rollback()
			http.Error(w, "El token ha expirado", http.StatusBadRequest)
			return
		}

		_, err = tx.Exec("UPDATE usuario SET activo = TRUE WHERE usuario_id = $1", userID)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al activar la cuenta", http.StatusInternalServerError)
			return
		}

		_, err = tx.Exec("DELETE FROM confirmacion_cuenta WHERE token = $1", token)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Error al eliminar el token de confirmación", http.StatusInternalServerError)
			return
		}

		if err := tx.Commit(); err != nil {
			http.Error(w, "Error al confirmar la transacción", http.StatusInternalServerError)
			return
		}*/

	http.Redirect(w, r, "/cuenta-activada", http.StatusSeeOther)
}
