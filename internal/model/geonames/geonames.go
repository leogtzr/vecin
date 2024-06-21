package model

// GeoNamesResponse ...
type GeoNamesResponse struct {
	Geonames []GeoName `json:"geonames"`
}

// GeoName ...
type GeoName struct {
	GeonameID   int    `json:"geonameId"`
	Name        string `json:"name"`
	ToponymName string `json:"toponymName"`
}
