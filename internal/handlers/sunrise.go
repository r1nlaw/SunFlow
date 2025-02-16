package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sunflow/internal/models"
)

func Handlers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Sunlight Map")
}

func SunlightHandler(w http.ResponseWriter, r *http.Request) {
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	if latitude == "" || longitude == "" {
		http.Error(w, "Should latitude and longitude not null", http.StatusBadRequest)
		return
	}

	sunlightData, err := models.GetSunLightData(latitude, longitude)
	if err != nil {
		http.Error(w, "Unable to fetch sunlight data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sunlightData)

}
