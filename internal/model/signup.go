package model

import "fmt"

type SignUpFormData struct {
	Username        string `json:"username"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Telefono        string `json:"telefono"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func (s SignUpFormData) String() string {
	return fmt.Sprintf("Nombre: (%s), Username: (%s), Apellido: (%s), Telefono: (%s), Email: (%s), Password: (%s), ConfirmPassword: (%s)",
		s.Nombre, s.Username, s.Apellido, s.Telefono, s.Email, s.Password, s.ConfirmPassword)
}
