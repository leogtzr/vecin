package model

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
