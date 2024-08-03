package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"vecin/internal/middleware"
	"vecin/internal/model"
	"vecin/internal/service"
)

func getUserIDFromSession(r *http.Request) (int, error) {
	session, err := middleware.GetSessionStore().Get(r, "session")
	if err != nil || session.Values["user_id"] == nil {
		return -1, err
	}

	userID := session.Values["user_id"].(int)

	return userID, nil
}

func DashboardPage(svc *service.Service, w http.ResponseWriter, r *http.Request) {
	// Check if the user is logged in:
	if loggedIn := isLoggedIn(r); !loggedIn {
		log.Print("debug:x DashboardPage, user is not logged in, redirecting to login page")
		// dev-note:bypass the following two lines to show the dashboard and check.
		//redirectLoginPage(w)
		//return
	}

	session, err := middleware.GetSessionStore().Get(r, "session")
	if err != nil || session.Values["user_id"] == nil {
		// If not logged in, redirect to the login page:
		redirectLoginPage(w)

		return
	}

	userID := session.Values["user_id"].(int)

	// 1) Verificar si el usuario ha registrado un fraccionamiento
	// 2) Verificar si el usuario está unido a un fraccionamiento
	registered, err := svc.ShouldShowWelcomePageIfNotRegistered(userID)
	if err != nil {
		log.Printf("error checking if the user is registered before showing welcome page: %v", err)
		redirectToWelcomePage(w)

		return
	}

	log.Printf("debug:x registered user is: %v", registered)

	if !registered {
		redirectToWelcomePage(w)

		return
	}

	log.Println("debug:x DashboardPage")

	redirectToDashboard(w, session)
}

// UpdateFracc updates a community by its ID.
// path: "/api/fraccionamientos/{communityID}"
func UpdateFracc(svc *service.Service, w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		writeMessageWithStatusCode(w, "method not allowed", http.StatusMethodNotAllowed)

		return
	}

	vars := mux.Vars(r)
	communityIDParam := vars["communityID"]
	communityID, err := strconv.Atoi(communityIDParam)
	if err != nil {
		writeMessageWithStatusCode(w, "invalid communityID", http.StatusBadRequest)

		return
	}

	log.Printf("debug:x fraccID=%d", communityID)

	var formData model.FraccionamientoFormData
	err = json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		redirectToErrorPageWithMessageAndStatusCode(w, "Unable to process input data", http.StatusInternalServerError)

		return
	}

	fmt.Println(formData)
	err = svc.UpdateFracc(communityID, formData)
	if err != nil {
		log.Printf("Error updating fraccion: %v", err)

		writeMessageWithStatusCode(w, "Error updating fracc", http.StatusInternalServerError)
		return
	}

	log.Printf("debug:x done updating fracc (communityID=(%d))", communityID)

	writeMessageWithStatusCode(w, "Successfully updated fracc", http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	store := middleware.GetSessionStore()
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Error obteniendo la sesión", http.StatusInternalServerError)
		return
	}

	// Verificar el token CSRF
	csrfToken := r.FormValue("csrf_token")
	if csrfToken != session.Values["csrf_token"] {
		http.Error(w, "Token CSRF inválido", http.StatusForbidden)
		return
	}

	// Clear all the values from session:
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Error guardando la sesión", http.StatusInternalServerError)
		return
	}

	//
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
