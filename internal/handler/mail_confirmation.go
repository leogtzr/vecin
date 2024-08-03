package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
	"vecin/internal/database"
	"vecin/internal/model"
	"vecin/internal/service"
)

func redirectAccountActivationProblem(w http.ResponseWriter) {
	templatePath := getTemplatePath("error-activate-account.html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)

	pageVariables := PageVariables{
		Year:     time.Now().Format("2006"),
		AppName:  "Vecin",
		LoggedIn: false,
	}

	err = t.Execute(w, pageVariables)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
}

func redirectAccountActivated(userConfirmedAccount model.Usuario, w http.ResponseWriter) {
	templatePath := getTemplatePath("account-activated.html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusSeeOther)

	pageVariables := struct {
		UserName string
	}{
		UserName: userConfirmedAccount.Username,
	}

	err = t.Execute(w, pageVariables)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
}

func ConfirmAccountHandler(svc *service.Service, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]

	log.Printf("debug:x token: (%s)", token)
	userConfirmedAccount, err := svc.ConfirmAccount(token)
	if err != nil {
		log.Printf("error: %v", err)
		redirectAccountActivationProblem(w)

		return
	}

	redirectAccountActivated(userConfirmedAccount, w)
}

func ConfirmAccountLinkSent(w http.ResponseWriter, r *http.Request) {
	templatePath := getTemplatePath("check-your-email-account.html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusSeeOther)

	pageVariables := PageVariables{
		Year:     time.Now().Format("2006"),
		AppName:  "Vecin",
		LoggedIn: false,
	}

	err = t.Execute(w, pageVariables)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
}

func ResendActivationEmail(dao *database.DAO, svc *service.Service, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeMessageWithStatusCode(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeMessageWithStatusCode(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := (*dao).GetUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeMessageWithStatusCode(w, "Correo no registrado", http.StatusNotFound)
			return
		}
		writeMessageWithStatusCode(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	if user.Activo {
		writeMessageWithStatusCode(w, "La cuenta ya está activada", http.StatusBadRequest)
		return
	}

	token, err := svc.CreateNewConfirmationAccountToken(user.ID)
	if err != nil {
		writeMessageWithStatusCode(w, "Error al enviar el correo de activación (token)", http.StatusInternalServerError)
		return
	}

	err = svc.SendConfirmationEmail(user.Username, user.Email, token)
	if err != nil {
		writeMessageWithStatusCode(w, "Error al enviar el correo de activación", http.StatusInternalServerError)
		return
	}

	writeMessageWithStatusCode(w, "Correo de activación reenviado exitosamente", http.StatusOK)
}
