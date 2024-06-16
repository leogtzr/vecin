package main

import (
	"log"
	"net/http"
	"os"
	"vecin/internal/dao"
	"vecin/internal/router"

	_ "github.com/lib/pq"
	"golang.org/x/time/rate"
)

var (
	dbMode     = os.Getenv("DB_MODE")
	dbHost     = os.Getenv("PGHOST")
	dbUser     = os.Getenv("PGUSER")
	dbPassword = os.Getenv("POSTGRES_PASSWORD")
	dbName     = os.Getenv("PGDATABASE")
	dbPort     = os.Getenv("PGPORT")
	runMode    = os.Getenv("RUN_MODE")
)

func init() {
	if dbMode == "" {
		dbMode = "postgres"
	}
	if runMode == "" {
		runMode = "dev"
	}
	if dbHost == "" {
		dbHost = os.Getenv("LEONLIB_DB_HOST")
		if dbHost == "" {
			dbHost = "localhost"
		}
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbUser == "" {
		dbUser = os.Getenv("VECIN_DB_USER")
	}
	if dbPassword == "" {
		dbPassword = os.Getenv("VECIN_DB_PASSWORD")
	}
	if dbName == "" {
		dbName = os.Getenv("VECIN_DB")
	}
}

func main() {
	log.Printf("DB mode: (%s)", dbMode)
	dao, err := dao.NewDAO(dbMode, dbHost, dbPort, dbUser, dbPassword, dbName)
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
	r := router.NewRouter(&dao, limiter)

	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8180"
	}

	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
