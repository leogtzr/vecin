package model

import (
	"fmt"
	"time"
)

type SignUpFormData struct {
	Username        string `json:"username"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Telefono        string `json:"telefono"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	PaymentOption   string `json:"paymentOption"`
}

// Al momento de guardar un usuario, necesitaremos un objeto con la información que hemos
// recibido del formulario desde el Front-end, también necesitaremos el Token generado
type SignUpUserConfig struct {
	SignUpFormData SignUpFormData
	Token          string
	ExpiryTime     time.Time
}

func (s SignUpFormData) String() string {
	return fmt.Sprintf("Nombre: (%s), Username: (%s), Apellido: (%s), Telefono: (%s), Email: (%s), Password: (%s), ConfirmPassword: (%s), PaymentOption: (%s)",
		s.Nombre, s.Username, s.Apellido, s.Telefono, s.Email, s.Password, s.ConfirmPassword, s.PaymentOption)
}
