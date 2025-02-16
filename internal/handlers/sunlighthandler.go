package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sunflow/internal/models" // Путь к вашей модели
	"time"
)

type SunlightResponse struct {
	Sunrise string `json:"sunrise"`
	Sunset  string `json:"sunset"`
	Daytime bool   `json:"daytime"`
}

// SunlightHandler - обработчик для получения данных о солнечном свете
func SunlightHandler(w http.ResponseWriter, r *http.Request) {
	// Для примера используем координаты Нью-Йорка (можно изменить для динамичных данных)
	latitude := 40.7128
	longitude := -74.0060

	// Получаем информацию о времени восхода и заката солнца
	data, err := models.GetSunlightData(latitude, longitude)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching sun data: %v", err), http.StatusInternalServerError)
		return
	}

	// Преобразуем время восхода и заката в объекты времени
	sunrise, err := time.Parse(time.RFC3339, data.Results.Sunrise)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing sunrise time: %v", err), http.StatusInternalServerError)
		return
	}
	sunset, err := time.Parse(time.RFC3339, data.Results.Sunset)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing sunset time: %v", err), http.StatusInternalServerError)
		return
	}

	// Получаем текущее время в UTC
	now := time.Now().UTC()

	// Определяем, день ли сейчас или ночь
	daytime := now.After(sunrise) && now.Before(sunset)

	// Создаем структуру с результатами
	response := SunlightResponse{
		Sunrise: sunrise.Format(time.Kitchen),
		Sunset:  sunset.Format(time.Kitchen),
		Daytime: daytime,
	}

	// Устанавливаем заголовки ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Отправляем результат в формате JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON response: %v", err), http.StatusInternalServerError)
		return
	}
}
