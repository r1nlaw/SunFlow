package handlers

import (
	"fmt"
	"net/http"
	"sunflow/internal/models"
)

func SunlightHandler(w http.ResponseWriter, r *http.Request) {
	// Например, широта и долгота для Нью-Йорка
	latitude := 40.7128
	longitude := -74.0060

	// Получаем информацию о времени восхода и заката солнца
	data, err := models.GetSunriseSunsetInfo(latitude, longitude)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching sun data: %v", err), http.StatusInternalServerError)
		return
	}

	// Выводим результат
	fmt.Fprintf(w, "Sunrise: %s\n", data.Results.Sunrise)
	fmt.Fprintf(w, "Sunset: %s\n", data.Results.Sunset)
}
