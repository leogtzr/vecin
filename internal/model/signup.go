package model

type SignUpFormData struct {
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Telefono        string `json:"modeloSuscripcion"`
	Email           string `json:"direccionCalle"`
	Password        string `json:"direccionNumero"`
	ConfirmPassword string `json:"direccionColonia"`
}

/*
nombre: $('#nombre').val(),
apellido: $('#apellido').val(),
telefono: $('#telefono').val(),
email: $('#email').val(),
password: $('#password').val(),
confirm_password: $('#confirm_password').val(),
*/
