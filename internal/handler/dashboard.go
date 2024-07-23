package handler

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"vecin/internal/service"
)

func DashboardPage(svc *service.Service, w http.ResponseWriter, r *http.Request) {
	// Check if the user is logged in:
	if loggedIn := isLoggedIn(r); !loggedIn {
		log.Print("debug:x DashboardPage, user is not logged in, redirecting to login page")
		//redirectLoginPage(w)
		//return
	}

	// Verificar si el usuario ha registrado un fraccionamiento
	// Verificar si el usuario está unido a un fraccionamiento

	log.Println("DashboardPage")

	pageVariables := PageVariables{
		Year:    time.Now().Format("2006"),
		AppName: "Vecin",
	}

	tmpl, err := template.ParseFiles(
		addTemplateFiles("internal/template/welcome.html")...,
	)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", pageVariables)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
