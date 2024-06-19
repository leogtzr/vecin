package model

import "time"

type Usuario struct {
	ID             int
	NombreUsuario  string
	NombreCompleto string
	Email          string
	HashContrasena string
	FechaCreacion  time.Time
}
