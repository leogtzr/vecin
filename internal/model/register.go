package model

type FraccionamientoFormData struct {
	NombreComunidad       string `json:"nombreComunidad"`
	TipoComunidad         string `json:"tipoComunidad"`
	ModeloSuscripcion     string `json:"modeloSuscripcion"`
	DireccionCalle        string `json:"direccionCalle"`
	DireccionNumero       string `json:"direccionNumero"`
	DireccionColonia      string `json:"direccionColonia"`
	DireccionCodigoPostal string `json:"direccionCodigoPostal"`
	DireccionCiudad       string `json:"direccionCiudad"`
	DireccionEstado       string `json:"direccionEstado"`
	DireccionPais         string `json:"direccionPais"`
	Referencias           string `json:"referencias"`
	Descripcion           string `json:"descripcion"`
}
