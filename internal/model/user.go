package model

import "time"

type Usuario struct {
	ID             int
	Username       string
	Nombre         string
	Apellido       string
	Telefono       string
	Email          string
	HashContrasena string
	Activo         bool
}

type UserExistence struct {
	Email  string
	Exists bool
}

// This model will help for the
type ConfirmationAccount struct {
	ConfirmationID int
	UserID         int
	Token          string
	ExpirationTime time.Time
}
