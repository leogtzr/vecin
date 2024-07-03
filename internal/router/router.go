package router

import (
	"net/http"
	"vecin/internal/config"
	"vecin/internal/database"
	"vecin/internal/handler"
	"vecin/internal/middleware"

	"golang.org/x/time/rate"

	"github.com/gorilla/mux"
)

type Router struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Router

func createRoutes(dao *database.DAO, cfg *config.Config) *Routes {
	routes := &Routes{
		Router{
			Name:        "IndexPage",
			Method:      "GET",
			Path:        "/",
			HandlerFunc: handler.IndexPage,
		},
		Router{
			Name:   "API - obtener los estados de un país",
			Method: "GET",
			Path:   "/api/region",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.GetRegionNameFromGeoNames(w, r, cfg)
			},
		},
		Router{
			Name:        "Landing Page",
			Method:      "GET",
			Path:        "/landing",
			HandlerFunc: handler.LandingPage,
		},
		Router{
			Name:        "Registrar Fraccionamiento",
			Method:      "GET",
			Path:        "/registrar-fraccionamiento",
			HandlerFunc: handler.RegisterFracc,
		},
		Router{
			Name:   "Form - Registrar Fraccionamiento",
			Method: "POST",
			Path:   "/registrar-fracc",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.FormRegisterFracc(dao, w, r)
			},
		},
		Router{
			Name:   "Registrar - Page",
			Method: "GET",
			Path:   "/registrar",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.RegisterPage(dao, w, r)
			},
		},
		Router{
			Name:   "View Fraccs",
			Method: "GET",
			Path:   "/view-fraccs",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.ViewFraccsPage(dao, w, r)
			},
		},
		Router{
			Name:   "Login",
			Method: "POST",
			Path:   "/login",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.Login(dao, w, r)
			},
		},
		Router{
			Name:   "Create Account - Page",
			Method: "GET",
			Path:   "/create-account",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.CreateAccountPage(dao, w, r)
			},
		},
		Router{
			Name:   "Create Account - Page",
			Method: "POST",
			Path:   "/create-account",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.SignUp(dao, w, r)
			},
		},
		Router{
			Name:   "Gen Error - Remove this",
			Method: "GET",
			Path:   "/generror",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.GenError(dao, w, r)
			},
		},
	}

	return routes
}

func NewRouter(dao *database.DAO, limiter *rate.Limiter, cfg *config.Config) *mux.Router {
	routes := createRoutes(dao, cfg)
	router := mux.NewRouter().StrictSlash(true)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	rateLimitMiddleware := middleware.RateLimitMiddlewareAdapter(limiter, nextHandler)
	router.Use(rateLimitMiddleware)

	for _, route := range *routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	fs := http.FileServer(http.Dir("assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	return router
}
