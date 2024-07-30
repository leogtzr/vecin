package model

// FraccionamientosResponse ...
type FraccionamientosResponse struct {
	Fraccionamientos []Fraccionamiento `json:"fraccionamientos"`
}

// Fraccionamiento ...
type Fraccionamiento struct {
	CommunityID       int    `json:"community_id"`
	Name              string `json:"name"`
	DireccionCalle    string `json:"direccion_calle"`
	DireccionNumero   string `json:"direccion_numero"`
	DireccionColonia  string `json:"direccion_colonia"`
	DireccionCP       string `json:"direccion_cp"`
	DireccionEstado   string `json:"direccion_estado"`
	DireccionCiudad   string `json:"direccion_ciudad"`
	DireccionPais     string `json:"direccion_pais"`
	ModeloSuscripcion string `json:"modelo_suscripcion"`
	Tipo              string `json:"tipo"`
	Referencias       string `json:"referencias"`
	Descripcion       string `json:"descripcion"`
}
