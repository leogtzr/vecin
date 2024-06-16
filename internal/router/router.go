package router

import (
	"net/http"
	"vecin/internal/dao"
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

func createRoutes(dao *dao.DAO) *Routes {
	routes := &Routes{
		Router{
			Name:        "IndexPage",
			Method:      "GET",
			Path:        "/",
			HandlerFunc: handler.IndexPage,
		},
	}

	return routes
}

func NewRouter(dao *dao.DAO, limiter *rate.Limiter) *mux.Router {
	routes := createRoutes(dao)
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
