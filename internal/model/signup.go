package model

import "fmt"

type SignUpFormData struct {
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Telefono        string `json:"telefono"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (s SignUpFormData) String() string {
	return fmt.Sprintf("Nombre: (%s), Apellido: (%s), Telefono: (%s), Email: (%s), Password: (%s), ConfirmPassword: (%s)",
		s.Nombre, s.Apellido, s.Telefono, s.Email, s.Password, s.ConfirmPassword)
}

/*
nombre: $('#nombre').val(),
apellido: $('#apellido').val(),
telefono: $('#telefono').val(),
email: $('#email').val(),
password: $('#password').val(),
confirm_password: $('#confirm_password').val(),
*/
