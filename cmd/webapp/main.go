package main

import (
	"log"
	"net/http"
	"os"
	"vecin/internal/config"
	"vecin/internal/database"
	"vecin/internal/router"

	_ "github.com/lib/pq"
	"golang.org/x/time/rate"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	dao, err := database.NewDAO(cfg.DBMode, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = dao.Close()
	}()

	err = dao.Ping()
	if err != nil {
		panic(err)
	}

	limiter := rate.NewLimiter(rate.Limit(1), 5)
	r := router.NewRouter(&dao, limiter, cfg)

	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8180"
	}

	log.Printf("Listening on port %s\n", cfg.HTTPPort)
	log.Fatal(http.ListenAndServe(":"+cfg.HTTPPort, r))
}
