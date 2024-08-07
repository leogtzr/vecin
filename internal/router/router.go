package router

import (
	"net/http"
	"vecin/internal/config"
	"vecin/internal/database"
	"vecin/internal/email"
	"vecin/internal/handler"
	"vecin/internal/middleware"
	"vecin/internal/service"

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

func createRoutes(svc *service.Service, dao *database.DAO, cfg *config.Config) *Routes {
	routes := &Routes{
		Router{
			Name:   "Dashboard Page",
			Method: "GET",
			Path:   "/dashboard",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.DashboardPage(svc, w, r)
			},
		},
		Router{
			Name:        "IndexPage",
			Method:      "GET",
			Path:        "/",
			HandlerFunc: handler.IndexPage,
		},
		Router{
			Name:   "API - obtener fraccionamientos/comunidades",
			Method: "GET",
			Path:   "/api/fraccionamientos",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.GetFraccionamientos(svc, w, r)
			},
		},
		Router{
			Name:   "API - obtener la información de un fraccionamiento",
			Method: "GET",
			Path:   "/api/fraccionamientos/{id}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.GetFraccionamientoByID(svc, w, r)
			},
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
			Name:   "Form - Actualizar Fraccionamiento",
			Method: "PUT",
			Path:   "/api/fraccionamientos/{communityID}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.UpdateFracc(svc, w, r)
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
			// TODO: redirigir finalmente a la página para ver los fraccionamientos registrados.
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.ViewFraccsPage(dao, w, r)
			},
		},
		Router{
			Name:   "Login Page",
			Method: "GET",
			Path:   "/login",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.LoginPage(dao, w, r)
			},
		},
		Router{
			Name:   "Perfil Page",
			Method: "GET",
			Path:   "/perfil",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.ProfilePage(svc, w, r)
			},
		},
		Router{
			Name:   "Logout",
			Method: "POST",
			Path:   "/logout",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.Logout(w, r)
			},
		},
		Router{
			Name:   "Login",
			Method: "POST",
			Path:   "/signIn",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.Login(dao, w, r)
			},
		},
		Router{
			Name:   "Check if email exists",
			Method: "POST",
			Path:   "/check-email",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.CheckEmail(svc, w, r)
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
				handler.SignUp(svc, w, r)
			},
		},
		Router{
			Name:   "Confirm Account",
			Method: "GET",
			Path:   "/confirmar-cuenta/{token}",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.ConfirmAccountHandler(svc, w, r)
			},
		},
		Router{
			Name:   "Confirm Account - waiting for confirmation through link",
			Method: "GET",
			Path:   "/confirm-account-pending",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.ConfirmAccountLinkSent(w, r)
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
		Router{
			Name:   "Resend Activation Link",
			Method: "POST",
			Path:   "/resend-activation",
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				handler.ResendActivationEmail(dao, svc, w, r)
			},
		},
	}

	return routes
}

func NewRouter(dao *database.DAO, limiter *rate.Limiter, cfg *config.Config) *mux.Router {
	emailSender := email.MailerSend{
		Config: cfg.Mailing,
	}
	svc := service.NewService(*dao, cfg, emailSender)
	routes := createRoutes(svc, dao, cfg)
	router := mux.NewRouter().StrictSlash(true)

	// TODO: check what to do with this handler...
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	if cfg.RateLimiterEnabled {
		rateLimitMiddleware := middleware.RateLimitMiddlewareAdapter(limiter, nextHandler)
		router.Use(rateLimitMiddleware)
	}

	for _, route := range *routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	fs := http.FileServer(http.Dir("assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	router.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	return router
}
