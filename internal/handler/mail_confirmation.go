package handler

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
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
