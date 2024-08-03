package database

import (
	"database/sql"
	"fmt"
	"vecin/internal/model"
)

type DAO interface {
	// Admin
	Close() error
	Ping() error

	// Usuarios
	GetUserByUsername(username string) (*model.Usuario, error)
	GetUserByEmail(email string) (*model.Usuario, error)
	UserExistsByEmail(email string) (bool, error)

	//Comunidades
	CreateCommunity(data model.FraccionamientoFormData, userID int) (int, error)
	UpdateCommunity(data model.FraccionamientoFormData, communityID int) (int, error)
	GetCommunitiesByUser(userID int) ([]model.Fraccionamiento, error)
	GetCommunityDetailsByID(id string) (model.Fraccionamiento, error)

	// Relaciones usuario-comunidad
	HasRegisteredAFracc(userID int) (bool, error)
	IsPartOfComunidad(userID int) (bool, error)

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
