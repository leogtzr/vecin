package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"vecin/internal/config"
	model "vecin/internal/model/geonames"
	"vecin/internal/service"
)

// GetRegionNameFromGeoNames calls GeoNames API to get a state or city information.
// path: "/fraccionamientos/region"
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

func GetFraccionamientos(svc *service.Service, w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromSession(r)
	if err != nil {
		log.Printf("Error al getUserIDFromSession for %d id: %v\n", userID, err)
		writeUnauthorized(w)
		return
	}

	fraccionamientos, err := svc.GetFraccionamientos(userID)
	if err != nil {
		log.Printf("Error al obtener fraccionamientos for %d id, error: %v\n", userID, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(fraccionamientos)
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ErrorResponse{Message: message}
	_ = json.NewEncoder(w).Encode(response)
}

func GetFraccionamientoByID(svc *service.Service, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	communityID := vars["id"]

	log.Printf("debug:x id: (%s)", communityID)
	fraccionamiento, err := svc.GetFraccionamientoDetail(communityID)
	if err != nil {
		log.Printf("Error al obtener detalles de fraccionamiento para el ID %d: %v\n", communityID, err)
		writeErrorResponse(w, http.StatusInternalServerError, "No se pudieron obtener los fraccionamientos")
		return
	}
	//userConfirmedAccount, err := svc.ConfirmAccount(token)
	//if err != nil {
	//	log.Printf("error: %v", err)
	//	redirectAccountActivationProblem(w)
	//
	//	return
	//}
	//
	//redirectAccountActivated(userConfirmedAccount, w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(fraccionamiento)
}
