package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"vecin/internal/config"
	model "vecin/internal/model/geonames"
)

// GetRegionNameFromGeoNames calls GeoNames API to get a state or city information.
// path: "/api/region"
func GetRegionNameFromGeoNames(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	geoNameId := r.URL.Query().Get("geonameId")

	log.Printf("Trying to get info for: %s", geoNameId)

	if geoNameId == "" {
		http.Error(w, "missing state param", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("http://api.geonames.org/childrenJSON?geonameId=%s&username=%s", geoNameId, cfg.GeoNamesUser)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error al realizar la solicitud HTTP: %v\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("error: closing body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error al leer la respuesta HTTP: %v\n", err)
		return
	}

	var geoNamesResponse model.GeoNamesResponse
	err = json.Unmarshal(body, &geoNamesResponse)
	if err != nil {
		log.Printf("Error al parsear JSON: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(geoNamesResponse)
}
