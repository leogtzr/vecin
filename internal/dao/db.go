package dao

import (
	"database/sql"
	"fmt"
)

// DAO
type DAO interface {
	Close() error
	Ping() error
}

type postgresBookDAO struct {
	db *sql.DB
}

func NewDAO(dbMode, dbHost, dbPort, dbUser, dbPassword, dbName string) (DAO, error) {
	var bookDAO DAO
	switch dbMode {
	case "postgres":
		var psqlInfo string

		psqlInfo = "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

		fmt.Printf("debug:x connection=(%s)\n", psqlInfo)

		DB, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return nil, err
		}
		bookDAO = &postgresBookDAO{
			db: DB,
		}

		return bookDAO, nil
	}

	return bookDAO, nil
}
