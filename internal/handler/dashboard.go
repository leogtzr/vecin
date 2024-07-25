package handler

import (
	"log"
	"net/http"
	"vecin/internal/middleware"
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

	redirectToDashboard(w)
}
