package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nadavbarlev/dream-exercise-go/internal/geo"
	"github.com/nadavbarlev/dream-exercise-go/internal/geo/fetcher"
	"github.com/nadavbarlev/dream-exercise-go/internal/models"
)

var geoFetcher fetcher.GeoFetcher = fetcher.NewIpApiGeoFetcher()
var geoResolver *geo.GeoResolver = geo.NewGeoResolver(geoFetcher)

func LocationHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	// Read the body JSON
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Decode the incoming body
	var locationRequest models.LocationRequest
	err = json.Unmarshal(body, &locationRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate location body fields
	if locationRequest.ReqID == "" || locationRequest.IP == "" {
    	http.Error(w, "ReqID and IP are required fields", http.StatusBadRequest)
        return
    }

	// Resolve the requested location 
	res, err := geoResolver.Resolve(locationRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert LocationResponse object to JSON
	jsonResponse, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}