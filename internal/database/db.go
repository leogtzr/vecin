package database

import (
	"database/sql"
	"fmt"
	"vecin/internal/model"
)

type DAO interface {
	Close() error
	Ping() error
	GetUserByUsername(username string) (*model.Usuario, error)
	SaveCommunity(data model.RegisterFormData) (int, error)
	DB() *sql.DB
}

type daoImpl struct {
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
		bookDAO = &daoImpl{
			db: DB,
		}

		return bookDAO, nil
	}

	return bookDAO, nil
}
